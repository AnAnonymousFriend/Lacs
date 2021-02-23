package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	log "Lacs/pkg/logging"
)

// @Summary 查找
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/login [Post]
func GetDevices(c *gin.Context){
	log.Error("Test Error")
	c.String(http.StatusOK, "hello, world")
}

