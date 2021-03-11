package routers

import (
	"Lacs/server/api"
	"github.com/gin-gonic/gin"
)

func InitUserCodeRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use()
	{
		UserRouter.GET("add", api.AddRole)
		UserRouter.GET("userAll", api.FindUserAll)
		UserRouter.POST("addUser",api.AddUser)
	}
}