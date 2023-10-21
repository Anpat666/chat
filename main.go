package main

import (
	"lottery/models"
	"lottery/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	models.InitMysql()
	routers.IndexRouter(r)
	r.Run()
}
