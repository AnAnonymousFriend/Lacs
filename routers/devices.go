package routers

import (
	"Lacs/server/api"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	DevicesRouter := Router.Group("autoCode").
		Use()
	{
		DevicesRouter.POST("Devices", api.GetDevices)

	}
}

