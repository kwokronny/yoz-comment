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
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	hour, min, _ := t.Clock()
	return fmt.Sprintf("%d-%02d-%02d %02d-%02d", year, month, day, hour, min)
}

// SetupRouter is setup router setting
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/index.html")
	engine.Use(middleware.LoggerToFile())

	if helper.Config.CROS_Enabled == true {
		engine.Use(cors.Default())
	}

	engine.GET("/page", comment.GetComment)
	engine.POST("/comment", comment.Save)

	manage := engine.Group(helper.Config.ManageRouter, gin.BasicAuth(gin.Accounts{
		helper.Config.AdminRoot: helper.Config.AdminPass,
	}))
	manage.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	manage.GET("/page", comment.GetPage)
	manage.POST("/delete", comment.Delete)

	return engine
}
