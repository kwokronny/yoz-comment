package main

import (
	"YozComment/router"
	"YozComment/util"
	"strconv"
)

func main() {
	r := router.SetupRouter()
	port := "9975"
	if util.Config.Installed == true {
		port = strconv.Itoa(util.Config.ServerPort)
	}
	r.Run(":" + port)
}
