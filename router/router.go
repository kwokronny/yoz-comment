package router

import (
	"YozComment/controller/comment"
	"YozComment/controller/manage"
	"YozComment/middleware"
	"YozComment/statics"
	"YozComment/util"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
)

// SetupRouter 装载路由
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	if util.Config.Installed == false {
		// engine.HTMLRender
		// engine.LoadHTMLFiles("templates/manage/install.html")
		engine.GET("/", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			tmpl, _ := statics.Asset("templates/manage/install.html")
			c.Writer.Write(tmpl)
			c.Writer.Flush()
			// c.HTML(http.StatusOK, "install.html", gin.H{})
		})
		engine.POST("/setting", util.SaveConfigFile)
	} else {
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

		// engine.LoadHTMLFiles("templates/web/index.html", "templates/manage/manage.html", "templates/manage/login.html")

		staticFS := assetfs.AssetFS{Asset: statics.Asset, AssetDir: statics.AssetDir, AssetInfo: statics.AssetInfo, Prefix: "templates/web/static", Fallback: "index.html"}
		engine.StaticFS("/static", &staticFS)
		// engine.StaticFile("/client.js", "./templates/web/client.js")

		engine.GET("/client.js", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			tmpl, _ := statics.Asset("templates/web/client.js")
			c.Writer.Write(tmpl)
			c.Writer.Header().Add("Content-Type", "application/javascript")
			c.Writer.Flush()
		})

		engine.GET("/index.html", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			tmpl, _ := statics.Asset("templates/web/index.html")
			c.Writer.Write(tmpl)
			c.Writer.Flush()
		})

		api := engine.Group("/api")
		api.GET("/page", comment.GetComment)
		api.POST("/comment", comment.Save)

		manageApi := engine.Group(util.Config.ManageRouter)
		manageApi.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, "./manage.html")
		})
		manageApi.GET("/login.html", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			tmpl, _ := statics.Asset("templates/manage/login.html")
			c.Writer.Write(tmpl)
			c.Writer.Flush()
		})

		manageApi.GET("/config.html", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			tmpl, _ := statics.Asset("templates/manage/install.html")
			c.Writer.Write(tmpl)
			c.Writer.Flush()
			// c.HTML(http.StatusOK, "install.html", gin.H{})
		})

		manageApi.POST("/getConfig", func(c *gin.Context) {
			util.Response{}.Success(c, util.Config)
		})

		manageApi.POST("/setting", util.SaveConfigFile)

		manageApi.GET("/manage.html", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			tmpl, _ := statics.Asset("templates/manage/manage.html")
			c.Writer.Write(tmpl)
			c.Writer.Flush()
		})
		manageApi.POST("/login", manage.Login)
		manageApi.GET("/page", middleware.AuthCheck(), manage.GetPage)
		manageApi.POST("/delete", middleware.AuthCheck(), manage.Delete)
	}
	return engine
}
