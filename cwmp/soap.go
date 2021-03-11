package cwmp

import (
	"encoding/xml"
	"fmt"
)

type XmlDevices struct {
	SoapHeader xml.Name `xml:"soap_env:Header"`
	SoapBody   xml.Name `xml:"soap_env:Body"`
	CwmpInform Inform
}

type Inform struct {
	XMLName    xml.Name `xml:"DeviceId"`
}

// SOPA为解析XML，和拼接CWMP协议报文的工具库

func analysisCwmp(rep []byte)  {
	v := XmlDevices{}
	_ = xml.Unmarshal(rep, &v)
	fmt.Println(v)
}