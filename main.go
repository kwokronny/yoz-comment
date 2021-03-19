package main

import (
	"YozComment/router"
	"YozComment/util"
	"strconv"
)

func main() {
	r := router.SetupRouter()
	r.Run(":" + strconv.Itoa(util.Config.ServerPort))
}
