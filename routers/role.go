package routers

import (
	"Lacs/server/api"
	"github.com/gin-gonic/gin"
)


func InitRoleCodeRouter(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("role").Use()
	{
		RoleRouter.POST("add", api.AddRole)
		RoleRouter.GET("findOne", api.FindOne)
		RoleRouter.GET("all", api.FindAll)
		RoleRouter.POST("deleteRole",api.DeleteRoleById)
	}
}