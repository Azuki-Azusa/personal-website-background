package server

import (
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

// Server runs.
func Server() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Delims("{{{", "}}}")
	router.LoadHTMLGlob("templates/*.html")

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/blog", func(c *gin.Context) {
		var param struct {
			Token    string
			Title    string
			Textarea string
			Content  string
		}
		c.Bind(&param)

		if param.Token != "orangestar" {
			c.JSON(200, gin.H{
				"msg": "token错误",
			})
			return
		}
		err := model.CommitNewBlog(param.Title, param.Textarea, param.Content)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": "提交失败",
			})
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"msg": "提交成功",
			})
		}
	})

	router.PUT("/blog", func(c *gin.Context) {
		var param struct {
			Token    string
			ID       int
			Title    string
			Textarea string
			Content  string
		}
		c.Bind(&param)

		if param.Token != "orangestar" {
			c.JSON(200, gin.H{
				"msg": "token错误",
			})
			return
		}
		err := model.EditBlog(param.ID, param.Title, param.Textarea, param.Content)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": "修改失败",
			})
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"msg": "修改成功",
			})
		}
	})

	router.GET("/bloglist", func(c *gin.Context) {
		token := c.Query("token")
		if token == "orangestar" {
			blogs, err := model.ReturnBlogList()
			if err != nil {
				c.JSON(500, gin.H{
					"msg": "查询失败",
				})
				panic(err)
			} else {
				c.JSON(200, gin.H{
					"bloglist": blogs,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"msg": "token错误",
			})
		}
	})

	router.POST("/lyric", func(c *gin.Context) {
		var param struct {
			Token    string
			Title    string
			Textarea string
		}
		c.Bind(&param)

		if param.Token != "orangestar" {
			c.JSON(200, gin.H{
				"msg": "token错误",
			})
			return
		}
		err := model.CommitNewLyric(param.Title, param.Textarea)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": "提交失败",
			})
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"msg": "提交成功",
			})
		}
	})

	router.PUT("/lyric", func(c *gin.Context) {
		var param struct {
			Token    string
			ID       int
			Title    string
			Textarea string
		}
		c.Bind(&param)

		if param.Token != "orangestar" {
			c.JSON(200, gin.H{
				"msg": "token错误",
			})
			return
		}
		err := model.EditLyric(param.ID, param.Title, param.Textarea)
		if err != nil {
			c.JSON(500, gin.H{
				"msg": "修改失败",
			})
			panic(err)
		} else {
			c.JSON(200, gin.H{
				"msg": "修改成功",
			})
		}
	})

	router.GET("/lyriclist", func(c *gin.Context) {
		token := c.Query("token")
		if token == "orangestar" {
			lyrics, err := model.ReturnLyricList()
			if err != nil {
				c.JSON(500, gin.H{
					"msg": "查询失败",
				})
				panic(err)
			} else {
				c.JSON(200, gin.H{
					"lyriclist": lyrics,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"msg": "token错误",
			})
		}
	})

	router.Run(":2434") // listen and serve on 0.0.0.0:2434

}
