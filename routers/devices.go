package routers

import (
	"Lacs/server/api"
	"github.com/gin-gonic/gin"
	"Lacs/middleware"
)


func InitAutoCodeRouter(Router *gin.RouterGroup) {
	DevicesRouter := Router.Group("Devices").Use(middleware.JWT())
	{
		DevicesRouter.GET("all", api.GetDevices)
	}
}

