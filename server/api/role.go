package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"Lacs/pkg/setting"
)

type role struct {
	ID string
	roleName string
}


// @Summary 角色
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/login [Get]
func AddRole(c *gin.Context) bool{
	roleName := c.PostForm("roleName")
	collection := setting.NewDataTableCollent(role)
	_, err := collection.InsertOne(context.Background(), roleName)
	if err !=nil {
		return false
	}
	return true
}
