package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)


type Beans struct {
	//XMLName xml.Name `xml:"beans"`
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ soap_env`
	RabbitQueues []RabbitQueue `xml:"http://www.springframework.org/schema/rabbit queue"`
	ListenerContainers []ListenerContainer `xml:"http://www.springframework.org/schema/rabbit listener-container"`
}

type RabbitQueue struct {
	Name string `xml:"name,attr"`
	Id string `xml:"id,attr"`
}

type ListenerContainer struct {
	Listeners []RabbitListener `xml:"http://www.springframework.org/schema/rabbit listener"`
}

type RabbitListener struct {
	Queues string `xml:"queues,attr"`
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

	fmt.Println(v)
	fmt.Println("SmtpServer : ",v.RabbitQueues)
	fmt.Println("SmtpServer : ",v.XMLName)

}