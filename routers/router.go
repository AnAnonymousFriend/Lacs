package routers

import (
	"github.com/gin-gonic/gin"
	//ginSwagger "github.com/swaggo/gin-swagger"
	"Lacs/middleware"
)

func Routers() *gin.Engine  {
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	Router.Use(middleware.CasbinHandler())
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ApiGroup := Router.Group("/api/v1")
	// 注册用户路由
	InitAutoCodeRouter(ApiGroup)
	InitRoleCodeRouter(ApiGroup)
	InitUserCodeRouter(ApiGroup)
	return Router
}