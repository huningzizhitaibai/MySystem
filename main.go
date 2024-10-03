package main

import (
	"github.com/gin-gonic/gin"
	"myTestProject/models"
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
			c.Redirect(302, "/home")
		} else {
			//登录失败
			//重新返回渲染登录界面
			c.Redirect(302, "/login")
		}
	})

	//处理用户注册
	MySystem.POST("/signup", func(c *gin.Context) {
		User := models.UserInfoBasic{
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
			Email:    c.PostForm("email"),
		}

		result := mysql.AddNewUser(User)
		if result {
			c.Redirect(302, "/login")
		} else {
			c.Redirect(302, "/signup")
		}
	})

	//渲染首页
	MySystem.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	//渲染提问的网页
	MySystem.GET("/ask", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ask.html", nil)
	})

	//提交新的问题
	MySystem.POST("/ask", func(c *gin.Context) {
		Question := models.Question{
			c.PostForm("title"),
			c.PostForm("tag"),
			c.PostForm("content"),
		}
		result := mysql.AddNewQuestion(Question)

		if result {
			c.Redirect(302, "/home")
		} else {
			c.Redirect(302, "/ask")
		}
	})

	//渲染搜索的页面
	MySystem.GET("/search", func(c *gin.Context) {
		c.HTML(http.StatusOK, "search.html", nil)
	})

	//提交搜索的关键词
	MySystem.POST("/search", func(c *gin.Context) {
		keyword := c.PostForm("keyword")
		qids := mysql.SearchInDatabase(keyword)
		for _, qid := range qids {
			question := mysql.ShowQuestionByQid(qid)
			c.JSON(200, gin.H{
				"title":   question.Title,
				"tag":     question.Tag,
				"content": question.Content,
			})
		}
	})

	MySystem.Run(":8080")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
