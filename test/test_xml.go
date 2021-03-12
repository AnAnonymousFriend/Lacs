package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)


type Beans struct {
	Header CwmpId `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header`
	ListenerContainers []ListenerContainer `xml:"http://www.springframework.org/schema/rabbit listener-container"`
}

type CwmpId struct {
	ID string `xml:"urn:dslforum-org:cwmp-1-2 ID"`
}




type ListenerContainer struct {
	Listeners []RabbitListener `xml:"http://www.springframework.org/schema/rabbit listener"`
}

type RabbitListener struct {
	Queues string `xml:"queues,attr"`
}

type CwmpID struct {
	ID xml.Name `xml:"urn:dslforum-org:cwmp-1-2 cwmp"`
}

type CwmpIDQ struct {
	ServerName string   `xml:"serverName"`
}

type RabbitQueue struct {
	Name string `xml:"name,attr"`
	Id string `xml:"id,attr"`
}

func main() {
	file, err := os.Open("test_sopa.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Beans{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v.Header.ID)
	fmt.Println(len(v.ListenerContainers))
	//fmt.Println(v.Header)
	//fmt.Println(v.Envelope)
	//fmt.Println("SmtpServer : ",v.Header)
	//fmt.Println("SmtpServer : ",v.HeaderTwo.ID)
	//fmt.Println("SmtpServer : ",v.XMLName)

}