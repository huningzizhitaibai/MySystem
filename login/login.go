package login

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//这个示例应该针对的是MongoDB数据库的，还需要改写适应mysql

type User struct {
	UserId   bson.ObjectId `bson:"_id,omitempty" json:"user_id"`
	UserName string
	Password string
}

// login data
type loginForm struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// login function
func login(c *gin.Context) {
	var form loginForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := findUserByUsername(form.UserName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}

	if user.Password != form.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
	}

	//能够运行到这里，说明用户名和密码已经核实正确，需要返回一个token保持连接的状态
	token, err := generateToken(user.UserId.Hex(), user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成临牌失败"})
		return
	}

	//运行到这里相当于才是所有功能都正常。
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func findUserByUsername(username string) (*User, error) {
	//这里是假装找到相关的用户并返回了一个结构体，包含这正确的userName和Password
	//在实际操作中，这里应该是访问相关的数据库，从而得到相关的数据应当改写
	return &User{UserName: username, Password: "password123"}, nil
}

// 这里token的作用是维持维持连接额状态，作用类似于cookie
func generateToken(UserId, UserName string) (string, error) {
	//这里也是简写了，假设生成了一串令牌
	//其实是要根据传进来的参数进行生成相关的token
	return "som_generate_token", nil
}

func main() {
	//默认注册了一个gin框架下的web实例，目前我是这么理解的
	router := gin.Default()

	//用到这个方法的时候访问的是/login的路由， 同时运行login函数，
	//但在这里我不是很理解，为什么login不用写成login(), 没有&应该也不是引用啊
	router.POST("/login", login)

	//将这个实例运行在8080端口
	router.Run(":8080")
}
