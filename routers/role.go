package routers

import (
	"Lacs/server/api"
	"github.com/gin-gonic/gin"
)


func InitRoleCodeRouter(Router *gin.RouterGroup) {
	DevicesRouter := Router.Group("role").Use()
	{
		DevicesRouter.GET("all", api.GetDevices)
	}
}