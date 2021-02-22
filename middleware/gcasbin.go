package middleware

import (
	"Lacs/pkg/setting"
	_ "Lacs/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
)



func CasbinVerification(sub string,dom string,obj string,act string)  {
	if passed, _ := setting.CabinEnforcer.Enforce(sub, dom, obj, act); passed {
		// permit clark to read data1
		fmt.Println("Enforce policy passed.")
	} else {
		// deny the request, show an error
		fmt.Println("Enforce policy denied.")
	}

}

func CasbinHandler() gin.HandlerFunc  {
	return func(c *gin.Context) {
		// 获取请求
		//claims, _ := c.Get("claims")
		//waitUse := claims.(*request.CustomClaims)
		//obj := c.Request.URL.RequestURI()

		//// 获取请求的URI
		//obj := c.Request.URL.RequestURI()
		//// 获取请求方法
		//act := c.Request.Method
		//// 获取用户的角色
		//sub := "AuthorityId"
		//
		//success, _ :=setting.CabinEnforcer.Enforce(sub, obj, act)
		//if success {
		//	c.Next()
		//} else {
		//	g := app.Gin{c}
		//	g.Response(http.StatusInternalServerError,e.ERROR_ACCESS_FORBIDDEN,nil)
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}

