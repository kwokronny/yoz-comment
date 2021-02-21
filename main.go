package main

import (
	"KBCommentAPI/helper"
	"KBCommentAPI/router"
	"strconv"
)

func main() {
	r := router.SetupRouter()
	r.Run(":" + strconv.Itoa(helper.Config.ServerPort))
}
