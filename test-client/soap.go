package main

import (
	"encoding/xml"
	"os"
)

type SoapEnvelope struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/wsdl/soap/ Envelope"`
	EncodingStyle string   `xml:"http://schemas.xmlsoap.org/wsdl/soap/ encodingStyle,attr"`
	Body          SoapBody `xml:"http://schemas.xmlsoap.org/wsdl/soap/ Body"`
}

type SoapBody struct {
	Body interface{}
}

func CreateSoapEnvelope() *SoapEnvelope {
	retval := &SoapEnvelope{}
	retval.EncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
	return retval
}

func main() {
	env := CreateSoapEnvelope()
	env.Body.Body = GetAllColumns{}

	encoder := xml.NewEncoder(os.Stdout)
	err := encoder.Encode(&env)
	if err != nil {
		println("Error encoding document:", err.Error())
		return
	}

}
