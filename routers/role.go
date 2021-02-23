package routers

import (
	"Lacs/server/api"
	"github.com/gin-gonic/gin"
)


func InitRoleCodeRouter(Router *gin.RouterGroup) {
	DevicesRouter := Router.Group("role").Use()
	{
		DevicesRouter.POST("add", api.AddRole)
	}
}