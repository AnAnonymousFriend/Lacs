package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 查找
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/login [Post]
func GetDevices(c *gin.Context){
	c.String(http.StatusOK, "hello, world")
}

