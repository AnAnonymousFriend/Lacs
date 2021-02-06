package main

import (
	"Lacs/server/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"

)

func main(){
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.POST("/devices",api.GetDevices)

}

func handleUDPConnection()  {
	listener,err := net.ListenUDP("udp",&net.UDPAddr{IP: net.ParseIP("127.0.0.1"),Port:9981})
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println("Local:<%s> \n",listener.LocalAddr().String())
	data := make([]byte,1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			fmt.Printf(err.Error())
		}


	}

	fmt.Println("Welcome to LACS！")
}