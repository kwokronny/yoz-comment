/*
 * @Author: KwokRonny
 * @Date: 2020-07-16 10:47:07
 * @LastEditors: KwokRonny
 * @LastEditTime: 2020-07-24 16:04:35
 * @Description: file content
 */
package router

import (
	"KBCommentAPI/controller/comment"
	"KBCommentAPI/helper"
	"KBCommentAPI/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter is setup router setting
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.LoadHTMLFiles("templates/static/index.html", "templates/manage.html")
	engine.Use(middleware.LoggerToFile())

	if helper.Config.CROSEnabled == true {
		engine.Use(cors.Default())
	}

	engine.Static("/static", "./templates/static")

	engine.GET("/", func(c *gin.Context) {
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
