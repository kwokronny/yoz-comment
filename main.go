package main

import (
	"KBCommentAPI/helper"
	"KBCommentAPI/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":" + helper.Config.ServerPort)
}
