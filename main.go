package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"myTestProject/AI"
	"myTestProject/middle"
	"myTestProject/models"
	"myTestProject/mysql"
	"net/http"
	"strconv"
)

func main() {
	//启动一个gin框架下的go服务
	MySystem := gin.Default()

	//为网页提供一个icon
	//MySystem.Use()

	//将所有的需要用户登录以后才能进行操作的接口放在这一路由下
	userGroup := MySystem.Group("/user")
	userGroup.Use(middle.JWTAuthMiddleware())
	{

		//渲染首页
		userGroup.GET("/home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.html", nil)
		})

		//渲染提问的网页
		userGroup.GET("/ask", func(c *gin.Context) {
			c.HTML(http.StatusOK, "ask.html", nil)
		})

		//提交新的问题
		userGroup.POST("/ask", func(c *gin.Context) {
			Question := models.Question{
				c.PostForm("title"),
				c.PostForm("tag"),
				c.PostForm("content"),
			}
			result := mysql.AddNewQuestion(Question)

			if result {
				c.JSON(http.StatusCreated, gin.H{
					"msg": "创建问题成功",
				})
				//c.Redirect(302, "/usr/home")
			} else {
				c.JSON(http.StatusExpectationFailed, gin.H{
					"msg": "添加失败请重试",
				})
				//c.Redirect(302, "/usr/ask")
			}
		})

		//渲染搜索的页面
		userGroup.GET("/search", func(c *gin.Context) {
			c.HTML(http.StatusOK, "search.html", nil)
		})

		//提交搜索的关键词
		userGroup.POST("/search", func(c *gin.Context) {
			keyword := c.PostForm("keyword")
			qids := mysql.SearchInDatabase(keyword)
			var Questons []models.Question
			for _, qid := range qids {
				question := mysql.ShowQuestionByQid(qid)
				Questons = append(Questons, question)
			}
			result, err := json.Marshal(Questons)
			if err != nil {
				c.JSON(http.StatusExpectationFailed, gin.H{})
			}
			c.Data(200, "application/json; charset=utf-8", result)
		})

		//渲染回答的页面
		userGroup.GET("/answer", func(c *gin.Context) {
			c.HTML(http.StatusOK, "answer.html", nil)
		})

		//提交回答的答案
		userGroup.POST("/answer", func(c *gin.Context) {
			userclaim, _ := middle.ParseToken(c.GetHeader("Authorization"))
			qid, _ := strconv.Atoi(c.PostForm("qid"))
			answer := models.Answer{
				c.PostForm("content"),
				mysql.CheckUsernameByID(int(userclaim.UserID)),
				//这里将回答相关问题的id值设成了表单传入，主要是不知道怎么合理设置路由了
				qid,
			}
			ret := mysql.AddNewAnswer(answer)
			if ret {
				c.JSON(http.StatusCreated, gin.H{
					"msg": "创建回答成功",
				})
			} else {
				c.JSON(http.StatusExpectationFailed, gin.H{
					"msg": "创建回答失败",
				})
			}
		})

		//提供特定的问题的先关数据
		//包括提问和它相关的回答
		//算了，还是只实现返回相关问题体吧，两种不同结构的数据不知道怎么处理了
		userGroup.GET("/question/:qid", func(c *gin.Context) {
			qid, _ := strconv.Atoi(c.Param("qid"))
			//暂时先不写了
			//question := mysql.GetQuestionByQid(qid)
			aids := mysql.SearchAidsByQid(qid)
			var Answers []models.Answer
			for _, aid := range aids {
				answer := mysql.GetAnswerByAid(aid)
				Answers = append(Answers, answer)
			}
			result, err := json.Marshal(Answers)
			if err != nil {
				c.JSON(http.StatusExpectationFailed, gin.H{})
			}
			c.Data(200, "application/json; charset=utf-8", result)

		})

		//调用大语言模型为用户提供回答
		userGroup.GET("/question/AI", func(c *gin.Context) {
			//这里使用Param方法是无法正常得到qid参数的，是为啥
			qid, _ := strconv.Atoi(c.Query("qid"))
			question := mysql.GetQuestionByQid(qid)

			answer := AI.AIapi(question.Content)
			//这个地方不能写python3估计是识别不出来，他也没做分类
			//cmd := exec.Command("python", "./AI/AIpic.py", question.Content)
			//out, _ := cmd.CombinedOutput()
			//src := string(out)
			//str := common.ReplaceForMe(src)
			//utf8Bytes := []byte(str)
			//
			//// 将字节切片转换为字符串
			//answer := string(utf8Bytes)

			// 打印字符串
			fmt.Println(answer)
			c.JSON(http.StatusOK, gin.H{
				"answer": answer,
			})
		})

		//删除相关
		userGroup.DELETE("/question/:qid", func(c *gin.Context) {})
	}
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
			token, _ := middle.GenerateToken(int64(mysql.GetUidByUsername(username)))
			c.JSON(200, gin.H{
				"token": token,
			})

			//下面的没成功，但好像也实现了接口
			//c.SetCookie("token", token, 3600, "/user/home", "", false, true)
			//
			//c.Redirect(http.StatusFound, "/user/home")
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

	MySystem.Run(":8080")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
