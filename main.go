package main

import (
	"chat/routers"
	"chat/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	r := gin.Default()
	routers.Index(r)
	r.Run()
}
