package main

import (
	"github.com/gin-gonic/gin"
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

	//填写登录数据
	MySystem.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//处理用户登录
	MySystem.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"msg":      "OK",
			"username": username,
			"password": password,
		})
	})

	MySystem.Run(":8080")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
