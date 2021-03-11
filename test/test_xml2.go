package main

import (
	"encoding/xml"
	"fmt"
)

var (
	str = `
   <?xml version="1.0" encoding="UTF-8" standalone="no"?>
<soap_env:Envelope
        xmlns:soap_env="http://schemas.xmlsoap.org/soap/envelope/"
        xmlns:soap_enc="http://schemas.xmlsoap.org/soap/encoding/"
        xmlns:xsd="http://www.w3.org/2001/XMLSchema"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xmlns:cwmp="urn:dslforum-org:cwmp-1-2"
        xmlns:rabbit="http://www.springframework.org/schema/rabbit"
        >
    <soap_env:Header>
        <cwmp:ID soap_env:mustUnderstand="1">62</cwmp:ID>
    </soap_env:Header>

    <rabbit:listener-container connection-factory="mqConnectionFactory" acknowledge="auto">
        <rabbit:listener queues="myQueue" ref="queueListenter"/>
    </rabbit:listener-container>
</soap_env:Envelope>
    `
)

type Response struct {
	XMLName xml.Name `xml:"Envelope"`
	AuthSuccess AuthSuccessStruct `xml:"Header"`
}

type AuthSuccessStruct struct {
	XMLName xml.Name `xml:"ID"`
	cwmp string `xml:"cwmp"`
	Ppal string `xml:"ppal"`
}


func main () {
	r := Response {}
	err := xml.Unmarshal([]byte(str), &r)
	fmt.Println(err, r)
	fmt.Printf("%v\n", r.AuthSuccess)
	fmt.Printf("user = %v\n", r.AuthSuccess.User)
	fmt.Printf("ppal = %v\n", r.AuthSuccess.Ppal)
}