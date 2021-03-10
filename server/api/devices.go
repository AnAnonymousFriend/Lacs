package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	log "Lacs/pkg/logging"
	s "Lacs/pkg/ssh"
)

// @Summary 查找
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/login [Post]
func GetDevices(c *gin.Context){
	log.Error("Test Error")
	c.String(http.StatusOK, "hello, world")
}

func DeviceCmd(c *gin.Context)  {
	oneDevice,_ := s.NewSwitchDevices("172.168.1.24",22,"admin","fs.com123")
	oneClient := &s.SwitchClient{
		Devices: oneDevice,
	}
	oneClient,_ =  oneClient.NewSShClient(oneDevice)
	cmds := []string{"show arp","show int status","exit","exit"}
	c.String(http.StatusOK, oneClient.DeviceCmd(cmds))
}

