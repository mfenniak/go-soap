package main

import (
	"encoding/xml"
	"os"
)

type WsdlDefinitions struct {
	Name            string          `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	TargetNamespace string          `xml:"http://schemas.xmlsoap.org/wsdl/ targetNamespace,attr"`
	Imports         []WsdlImport    `xml:"http://schemas.xmlsoap.org/wsdl/ import"`
	Types           []WsdlTypes     `xml:"http://schemas.xmlsoap.org/wsdl/ types"`
	Messages        []WsdlMessage   `xml:"http://schemas.xmlsoap.org/wsdl/ message"`
	PortTypes       []WsdlPortTypes `xml:"http://schemas.xmlsoap.org/wsdl/ portType"`
	Services        []WsdlService   `xml:"http://schemas.xmlsoap.org/wsdl/ service"`
	Bindings        []WsdlBinding   `xml:"http://schemas.xmlsoap.org/wsdl/ binding"`
}

type WsdlBinding struct {
	Name         string          `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Type         string          `xml:"http://schemas.xmlsoap.org/wsdl/ type,attr"`
	Operations   []WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
	SoapBindings []SoapBinding   `xml:"http://schemas.xmlsoap.org/wsdl/soap/ binding"`
}

type SoapBinding struct {
	Transport string `xml:"http://schemas.xmlsoap.org/wsdl/soap/ transport,attr"`
}

type WsdlTypes struct {
	XsdSchema []XsdSchema `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

type WsdlImport struct {
	Namespace string `xml:"http://schemas.xmlsoap.org/wsdl/ namespace,attr"`
	Location  string `xml:"http://schemas.xmlsoap.org/wsdl/ location,attr"`
}

type WsdlMessage struct {
	Name  string            `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Parts []WsdlMessagePart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

type WsdlMessagePart struct {
	Name    string `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Element string `xml:"http://schemas.xmlsoap.org/wsdl/ element,attr"`
}

type WsdlPortTypes struct {
	Name       string          `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Operations []WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

type WsdlOperation struct {
	Name           string                `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Inputs         []WsdlOperationInput  `xml:"http://schemas.xmlsoap.org/wsdl/ input"`
	Outputs        []WsdlOperationOutput `xml:"http://schemas.xmlsoap.org/wsdl/ output"`
	Faults         []WsdlOperationFault  `xml:"http://schemas.xmlsoap.org/wsdl/ fault"`
	SoapOperations []SoapOperation       `xml:"http://schemas.xmlsoap.org/wsdl/soap/ operation"`
}

type WsdlOperationInput struct {
	Message    string `xml:"http://schemas.xmlsoap.org/wsdl/ message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlOperationOutput struct {
	Message    string `xml:"http://schemas.xmlsoap.org/wsdl/ message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlOperationFault struct {
	Name       string `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Message    string `xml:"http://schemas.xmlsoap.org/wsdl/ message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlService struct {
	Name  string     `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Ports []WsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}

type WsdlPort struct {
	Name          string        `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Binding       string        `xml:"http://schemas.xmlsoap.org/wsdl/ binding,attr"`
	SoapAddresses []SoapAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
}

type SoapAddress struct {
	Location string `xml:"http://schemas.xmlsoap.org/wsdl/soap/ location,attr"`
}

type SoapOperation struct {
	SoapAction string `xml:"http://schemas.xmlsoap.org/wsdl/soap/ soapAction,attr"`
	Style      string `xml:"http://schemas.xmlsoap.org/wsdl/soap/ style,attr"`
}

type SoapBody struct {
	Use string `xml:"http://schemas.xmlsoap.org/wsdl/soap/ use,attr"`
}

type XsdSchema struct {
	TargetNamespace string      `xml:"http://www.w3.org/2001/XMLSchema targetNamespace,attr"`
	Imports         []XsdImport `xml:"http://www.w3.org/2001/XMLSchema import"`
}

type XsdImport struct {
	SchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema schemaLocation,attr"`
	Namespace      string `xml:"http://www.w3.org/2001/XMLSchema namespace,attr"`
}

func main() {
	//r, err := os.Open("ProjectListService1.svc?wsdl")
	r, err := os.Open("ProjectListService1.svc?wsdl=wsdl0")
	if err != nil {
		println("Error opening test.xml:", err.Error())
		return
	}

	decoder := xml.NewDecoder(r)
	obj := WsdlDefinitions{}
	err = decoder.Decode(&obj)
	if err != nil {
		println("Error decoding document:", err.Error())
		return
	}

	println("WSDL:", obj.TargetNamespace, obj.Name)
	for _, imp := range obj.Imports {
		println("Import:", imp.Namespace, imp.Location)
	}
	for _, types := range obj.Types {
		for _, schema := range types.XsdSchema {
			println("Schema:", schema.TargetNamespace)
			for _, imp := range schema.Imports {
				println("\tImport:", imp.SchemaLocation, imp.Namespace)
			}
		}
	}
	for _, message := range obj.Messages {
		println("Message:", message.Name)
		for _, part := range message.Parts {
			println("\tPart:", part.Name, part.Element)
		}
	}
	for _, portType := range obj.PortTypes {
		println("PortType:", portType.Name)
		for _, operation := range portType.Operations {
			println("\tOperation:", operation.Name)
			for _, input := range operation.Inputs {
				println("\t\tInput:", input.Message, input.WsawAction)
			}
			for _, output := range operation.Outputs {
				println("\t\tOutput:", output.Message, output.WsawAction)
			}
			for _, fault := range operation.Faults {
				println("\t\tFault:", fault.Message, fault.WsawAction)
			}
		}
	}
	for _, service := range obj.Services {
		println("Service:", service.Name)
		for _, port := range service.Ports {
			println("\tPort:", port.Name, port.Binding)
			for _, address := range port.SoapAddresses {
				println("\t\tAddress:", address.Location)
			}
		}
	}
	for _, binding := range obj.Bindings {
		println("Binding:", binding.Name, binding.Type)
		for _, soapBinding := range binding.SoapBindings {
			println("\tSOAP Binding:", soapBinding.Transport)
		}
		for _, operation := range binding.Operations {
			println("\tOperation:", operation.Name)
			for _, soapOperation := range operation.SoapOperations {
				println("\t\tSOAP Operation:", soapOperation.SoapAction, soapOperation.Style)
			}
		}
	}

	/*
		w, err := os.OpenFile("OutTest.wsdl", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			println("Error opening OutTest.wsdl:", err.Error())
			return
		}

		_, err = w.WriteString(xml.Header)
		if err != nil {
			println("Error writing header:", err.Error())
			return
		}

		encoder := xml.NewEncoder(w)
		err = encoder.Encode(&obj)
		if err != nil {
			println("Error encoding document:", err.Error())
			return
		}

		w.Close()
	*/
}
