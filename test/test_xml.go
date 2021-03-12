package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)


type Beans struct {
	Header CwmpId `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header`
	Body XmlBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body`
	ListenerContainers []ListenerContainer `xml:"http://www.springframework.org/schema/rabbit listener-container"`
}

type CwmpId struct {
	ID string `xml:"urn:dslforum-org:cwmp-1-2 ID"`
}

type XmlBody struct {
	Inform DeviceInform `xml:"urn:dslforum-org:cwmp-1-2 Inform"`
}

type DeviceInform struct {
	DeviceId     DeviceInfo   `xml:"DeviceId"`
	Event		Event	`xml:"Event"`

	MaxEnvelopes int `xml:"MaxEnvelopes"`
	CurrentTime string `xml:"CurrentTime"`
	RetryCount int `xml:"RetryCount"`

	ParameterList ParameterList `xml:"ParameterList"`
}

type  ParameterList struct {
	ParameterValueStruct []ParameterValueStruct `xml:"ParameterValueStruct"`
}

type ParameterValueStruct struct {
	Name string `xml:"Name"`
	Value string `xml:"Value"`
}

type  Event struct {
	EventStruct []EventStruct `xml:"EventStruct"`
}

type  EventStruct struct {
	EventCode string `xml:"EventCode"`
}

type DeviceInfo struct {
	Manufacturer     string   `xml:"Manufacturer"`
	OUI     		 string   `xml:"OUI"`
	ProductClass     string   `xml:"ProductClass"`
	SerialNumber     string   `xml:"SerialNumber"`
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
	fmt.Println(v.Body.Inform.DeviceId.Manufacturer)
	fmt.Println(v.Body.Inform.Event.EventStruct[0].EventCode)
	fmt.Println(v.Body.Inform.MaxEnvelopes)
	fmt.Println(v.Body.Inform.CurrentTime)

	fmt.Println(v.Body.Inform.ParameterList.ParameterValueStruct[0].Name)
	fmt.Println(v.Body.Inform.ParameterList.ParameterValueStruct[0].Value)
}