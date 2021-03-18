package main

import (
	"YozComment/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter is setup router setting
func main() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/install.html")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "install.html", gin.H{})
	})
	engine.POST("/setting", util.SaveConfigFile)

	engine.Run(":8080")
}
