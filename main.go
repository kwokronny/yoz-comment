package main

import (
	"YozComment/helper"
	"YozComment/router"
	"strconv"
)

func main() {
	r := router.SetupRouter()
	r.Run(":" + strconv.Itoa(util.Config.ServerPort))
}
