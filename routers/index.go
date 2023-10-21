package routers

import (
	"lottery/controllers/index"

	"github.com/gin-gonic/gin"
)

func IndexRouter(r *gin.Engine) {
	indexrouter := r.Group("/")
	{
		indexrouter.GET("/", index.IndexController{}.Index)

		indexrouter.GET("/index", index.IndexController{}.IndexOk)

		indexrouter.POST("/login", index.PlayerController{}.Login)

		indexrouter.GET("/register", index.PlayerController{}.Register)

		indexrouter.POST("/register/confirm", index.PlayerController{}.RegisterConfirm)

	}
}
