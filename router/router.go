/*
 * @Author: KwokRonny
 * @Date: 2020-07-16 10:47:07
 * @LastEditors: KwokRonny
 * @LastEditTime: 2020-07-24 16:04:35
 * @Description: file content
 */
package router

import (
	"fmt"
	"html/template"
	"kwok-comment/controller/comment"
	"kwok-comment/middleware"
	"time"

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
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	engine.LoadHTMLGlob("templates/index.tmpl")
	engine.Use(middleware.LoggerToFile())

	r := engine.Group("api-comment")

	r.GET("/page", comment.GetComment)
	r.POST("/comment", comment.Save)

	manage := r.Group("manage", gin.BasicAuth(gin.Accounts{
		"Hwq0za9WlCNSS7pT": "h#EqOHOr#Yl&v)ah",
	}))
	manage.POST("/delete", comment.Delete)
	manage.GET("/page", comment.GetPage)

	return engine
}
