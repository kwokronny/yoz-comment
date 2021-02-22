package router

import (
	"KBCommentAPI/controller/comment"
	"KBCommentAPI/helper"
	"KBCommentAPI/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter is setup router setting
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.LoadHTMLFiles("templates/static/index.html", "templates/manage.html")
	engine.Use(middleware.LoggerToFile())

	if helper.Config.CROSEnabled == true {
		engine.Use(func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}

			c.Next()
		})
	}

	engine.Static("/static", "./templates/static")
	engine.StaticFile("client.js", "./templates/static/client.js")

	engine.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	api := engine.Group("/api")
	api.GET("/page", comment.GetComment)
	api.POST("/comment", comment.Save)

	manage := engine.Group(helper.Config.ManageRouter, gin.BasicAuth(gin.Accounts{
		helper.Config.AdminRoot: helper.Config.AdminPass,
	}))
	manage.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "manage.html", gin.H{})
	})
	manage.GET("/page", comment.GetPage)
	manage.POST("/delete", comment.Delete)

	return engine
}
