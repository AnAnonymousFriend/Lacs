package cwmp

import (
	"encoding/xml"
	"fmt"
)


type DevicesXml struct {
	Header CwmpId `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header`
	Body XmlBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body`
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

// SOPA为解析XML，和拼接CWMP协议报文的工具库

func analysisCwmp(rep []byte)  {
	v := DevicesXml{}
	_ = xml.Unmarshal(rep, &v)

	fmt.Println(v.Header.ID)
	fmt.Println(v.Body.Inform.DeviceId.Manufacturer)
	fmt.Println(v.Body.Inform.Event.EventStruct[0].EventCode)
	fmt.Println(v.Body.Inform.MaxEnvelopes)
	fmt.Println(v.Body.Inform.CurrentTime)
	fmt.Println(v.Body.Inform.ParameterList.ParameterValueStruct[0].Name)
	fmt.Println(v.Body.Inform.ParameterList.ParameterValueStruct[0].Value)
}