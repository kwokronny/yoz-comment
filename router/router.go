package router

import (
	"YozComment/controller/comment"
	"YozComment/controller/manage"
	"YozComment/middleware"
	"YozComment/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter is setup router setting
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	if util.Config.Installed == false {
		engine.LoadHTMLFiles("templates/manage/install.html")
		engine.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "install.html", gin.H{})
		})
		engine.POST("/setting", util.SaveConfigFile)
	} else {
		engine.LoadHTMLFiles("templates/web/index.html", "templates/manage/manage.html", "templates/manage/login.html", "templates/install/login.html")
		engine.Use(middleware.LoggerMiddleware())

		if util.Config.CROSEnabled == true {
			engine.Use(func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

				if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
				}

				c.Next()
			})
		}

		engine.Static("/static", "./templates/web/static")
		engine.StaticFile("client.js", "./templates/web/client.js")

		engine.GET("/index.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})

		api := engine.Group("/api")
		api.GET("/page", comment.GetComment)
		api.POST("/comment", comment.Save)

		manageApi := engine.Group(util.Config.ManageRouter, func(c *gin.Context) {
			c.Redirect(http.StatusPermanentRedirect, "/index.html")
		})
		manageApi.GET("/login.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		})
		manageApi.GET("/index.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "manage.html", gin.H{})
		})
		manageApi.POST("/login", manage.Login)
		manageApi.GET("/page", middleware.AuthCheck(), manage.GetPage)
		manageApi.POST("/delete", middleware.AuthCheck(), manage.Delete)

	}
	return engine
}
