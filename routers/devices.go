package routers

import (
	"Lacs/server/api"
	"github.com/gin-gonic/gin"
)


func InitAutoCodeRouter(Router *gin.RouterGroup) {
	DevicesRouter := Router.Group("Devices").Use()
	{
		DevicesRouter.GET("all", api.GetDevices)
		DevicesRouter.GET("cmd", api.DeviceCmd)
	}
}

