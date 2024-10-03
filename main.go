package main

import (
	"github.com/gin-gonic/gin"
	"myTestProject/mysql"
	"net/http"
)

func main() {
	//启动一个gin框架下的go服务
	MySystem := gin.Default()

	//为网页提供一个icon
	//MySystem.Use()

	//加载静态页面,但是我理解的这里可能只是将所有的页面进行了缓存
	MySystem.LoadHTMLGlob("template/*.html")

	//在前端相应一个页面
	MySystem.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//渲染填写登录数据的页面
	MySystem.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//渲染注册的页面
	MySystem.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	//处理用户登录
	MySystem.POST("/login", func(c *gin.Context) {
		//从这里获得了，之前用户填写的相关数据
		username := c.PostForm("username")
		password := c.PostForm("password")

		//对用户的账户进行登录的验证
		result := mysql.CheckUser(username, password)
		if result {
			c.JSON(http.StatusOK, gin.H{
				"msg":      "登录成功",
				"username": username,
			})
		} else {
			//登录失败
			c.JSON(http.StatusForbidden, gin.H{})
		}
	})

	//处理用户注册
	MySystem.POST("/signup", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		//comfirm := c.PostForm("comfirm")
		//email := c.PostForm("email")
		result := mysql.AddNewUser(username, password)
		if result {
			c.JSON(http.StatusCreated, gin.H{
				"msg": "注册成功",
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "注册失败",
			})
		}
	})
	MySystem.Run(":8080")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
