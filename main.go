package main

import (
	"Lacs/pkg/setting"

	"Lacs/routers"
	_ "Lacs/server/api"
	"fmt"
	"net"

)

func main(){
	setting.Setup()
	//setting.ScheduleSetup()
	ginRouter := routers.Routers()
	ginRouter.Run(setting.AppSetting.Host)

}

func handleUDPConnection()  {
	listener,err := net.ListenUDP("udp",&net.UDPAddr{IP: net.ParseIP("127.0.0.1"),Port:9981})
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println("Local :<%s>\n",listener.LocalAddr().String())
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