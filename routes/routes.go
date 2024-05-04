package routes

import (
	"bing16Service/internal/service/blog"
	"bing16Service/internal/service/sslx"
	"bing16Service/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleReq struct {
	Context string `json:"context"`
}

type SslxReq struct {
	Data string `json:"data"`
}

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/blog/wyzx/:name", func(c *gin.Context) {
		var newBlog blog.Blog
		newBlog.Path = "wyzx"
		newBlog.Title = c.Param("name")

		var article ArticleReq
		if err := c.BindJSON(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newBlog.Context = article.Context
		blog.AddWyzx(newBlog)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Add success: " + newBlog.Title,
		})
	})

	r.POST("/blog/sslx", func(c *gin.Context) {
		var sslxReq SslxReq
		if err := c.BindJSON(&sslxReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := sslx.Add(sslxReq.Data)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "fail",
				"message": "Add fail: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Add success",
		})
	})
	return r
}
