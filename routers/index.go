package routers

import (
	"chat/docs"
	"chat/service"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/index", service.GetIndex)
	r.GET("/user", service.GetUserList)
	r.GET("/user/create", service.UserCreate)
	r.GET("/user/delete", service.UserDelete)
	r.POST("/user/updata", service.UserUpdata)
	r.POST("/user/login", service.UserLogin)
	r.POST("/user/changepwd", service.UserChangePassword)

	r.GET("/user/sendmsg", service.SendMsg)
}
