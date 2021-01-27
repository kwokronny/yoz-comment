package main

import (
	"kwok-comment/helper"
	"kwok-comment/router"
)

func main() {
	r := router.SetupRouter()

	r.Run(":" + helper.Config.ServerPort)
}
