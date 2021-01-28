/*
 * @Author: KwokRonny
 * @Date: 2020-07-16 10:47:07
 * @LastEditors: KwokRonny
 * @LastEditTime: 2020-07-24 16:04:35
 * @Description: file content
 */
package router

import (
	"kwok-comment/controller/comment"
	"kwok-comment/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter is setup router setting
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.Use(middleware.LoggerToFile())

	r := engine.Group("api-comment")

	r.GET("/page", comment.GetPage)
	r.POST("/comment", comment.Save)

	return engine
}
