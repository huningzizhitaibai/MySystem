package middle

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
这里是事先确定好一个最终答案
通过加密，将这个答案与用户id进行计算获得token
逆向时，需要调用用户ID反向进行计算获得jwtKey，同时满足鉴权和将token和用户进行绑定
*/
var jwtKey = []byte("key")

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 创建Token
// 根据用户的ID进行创建
func GenerateToken(userID int64) (string, error) {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			//为这个claim赋一个可用的时间
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),

			//设置签发人，但我不知道有什么用目前
			Issuer: "huning",
		},
	}

	//通过jwt包中的方法通过之前设置的信息，通过一定的加密算法，计算出一个token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//将这个token转化成string类型返回，方便使用
	return token.SignedString(jwtKey)
}

// ParseToken 反向解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	//将解析的token中的元素归类，用Claim方法进行赋值
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// JWT 中间件主体
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//如果有写token的先关的话，浏览器会自动记录token，在下次发送请求的时候
		//直接记录在request报文中
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}
		claims, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		//这set有什么用？
		c.Set("claims", claims)
		c.Next()
	}
}
