package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"strings"
)

type WsdlDefinitions struct {
	Name            string           `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	TargetNamespace string           `xml:"http://schemas.xmlsoap.org/wsdl/ targetNamespace,attr"`
	Imports         []*WsdlImport    `xml:"http://schemas.xmlsoap.org/wsdl/ import"`
	Types           []*WsdlTypes     `xml:"http://schemas.xmlsoap.org/wsdl/ types"`
	Messages        []*WsdlMessage   `xml:"http://schemas.xmlsoap.org/wsdl/ message"`
	PortTypes       []*WsdlPortTypes `xml:"http://schemas.xmlsoap.org/wsdl/ portType"`
	Services        []*WsdlService   `xml:"http://schemas.xmlsoap.org/wsdl/ service"`
	Bindings        []*WsdlBinding   `xml:"http://schemas.xmlsoap.org/wsdl/ binding"`
}

type WsdlBinding struct {
	Name         string           `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Type         string           `xml:"http://schemas.xmlsoap.org/wsdl/ type,attr"`
	Operations   []*WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
	SoapBindings []*SoapBinding   `xml:"http://schemas.xmlsoap.org/wsdl/soap/ binding"`
}

type SoapBinding struct {
	Transport string `xml:"http://schemas.xmlsoap.org/wsdl/soap/ transport,attr"`
}

type WsdlTypes struct {
	XsdSchema []*XsdSchema `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

type WsdlImport struct {
	Namespace string `xml:"http://schemas.xmlsoap.org/wsdl/ namespace,attr"`
	Location  string `xml:"http://schemas.xmlsoap.org/wsdl/ location,attr"`
}

type WsdlMessage struct {
	Name  string             `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Parts []*WsdlMessagePart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

type WsdlMessagePart struct {
	Name    string `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Element string `xml:"http://schemas.xmlsoap.org/wsdl/ element,attr"`
}

type WsdlPortTypes struct {
	Name       string           `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Operations []*WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

type WsdlOperation struct {
	Name           string                 `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Inputs         []*WsdlOperationInput  `xml:"http://schemas.xmlsoap.org/wsdl/ input"`
	Outputs        []*WsdlOperationOutput `xml:"http://schemas.xmlsoap.org/wsdl/ output"`
	Faults         []*WsdlOperationFault  `xml:"http://schemas.xmlsoap.org/wsdl/ fault"`
	SoapOperations []*SoapOperation       `xml:"http://schemas.xmlsoap.org/wsdl/soap/ operation"`
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
	Name  string      `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Ports []*WsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}

type WsdlPort struct {
	Name          string         `xml:"http://schemas.xmlsoap.org/wsdl/ name,attr"`
	Binding       string         `xml:"http://schemas.xmlsoap.org/wsdl/ binding,attr"`
	SoapAddresses []*SoapAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
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
	TargetNamespace    string            `xml:"http://www.w3.org/2001/XMLSchema targetNamespace,attr"`
	ElementFormDefault string            `xml:"http://www.w3.org/2001/XMLSchema elementFormDefault,attr"`
	Imports            []*XsdImport      `xml:"http://www.w3.org/2001/XMLSchema import"`
	Elements           []*XsdElement     `xml:"http://www.w3.org/2001/XMLSchema element"`
	ComplexTypes       []*XsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type XsdImport struct {
	SchemaLocation string `xml:"http://www.w3.org/2001/XMLSchema schemaLocation,attr"`
	Namespace      string `xml:"http://www.w3.org/2001/XMLSchema namespace,attr"`
}

type XsdElement struct {
	Name        string          `xml:"http://www.w3.org/2001/XMLSchema name,attr"`
	Nillable    bool            `xml:"http://www.w3.org/2001/XMLSchema nillable,attr"`
	Type        string          `xml:"http://www.w3.org/2001/XMLSchema type,attr"`
	MinOccurs   string          `xml:"http://www.w3.org/2001/XMLSchema minOccurs,attr"`
	MaxOccurs   string          `xml:"http://www.w3.org/2001/XMLSchema maxOccurs,attr"`
	ComplexType *XsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	SimpleType  *XsdSimpleType  `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

type XsdComplexType struct {
	Name     string       `xml:"http://www.w3.org/2001/XMLSchema name,attr"`
	Sequence *XsdSequence `xml:"http://www.w3.org/2001/XMLSchema sequence"`
}

type XsdSimpleType struct {
	Name     string          `xml:"http://www.w3.org/2001/XMLSchema name,attr"`
	Sequence *XsdRestriction `xml:"http://www.w3.org/2001/XMLSchema restriction"`
}

type XsdSequence struct {
	Elements []*XsdElement `xml:"http://www.w3.org/2001/XMLSchema element"`
}

type XsdRestriction struct {
	Base         string           `xml:"http://www.w3.org/2001/XMLSchema base,attr"`
	Pattern      *XsdPattern      `xml:"http://www.w3.org/2001/XMLSchema pattern"`
	MinInclusive *XsdMinInclusive `xml:"http://www.w3.org/2001/XMLSchema minInclusive"`
	MaxInclusive *XsdMaxInclusive `xml:"http://www.w3.org/2001/XMLSchema maxInclusive"`
}

type XsdPattern struct {
	Value string `xml:"http://www.w3.org/2001/XMLSchema value,attr"`
}

type XsdMinInclusive struct {
	Value string `xml:"http://www.w3.org/2001/XMLSchema value,attr"`
}

type XsdMaxInclusive struct {
	Value string `xml:"http://www.w3.org/2001/XMLSchema value,attr"`
}

func main() {

	wsdlMap := make(map[string]*WsdlDefinitions)
	xsdMap := make(map[string]*XsdSchema)

	_, err := LoadWsdl("http://na2.replicon.com/services/ProjectService1.svc?wsdl", wsdlMap, xsdMap)
	if err != nil {
		println("Error fetching WSDL:", err.Error())
		return
	}

	w, err := os.OpenFile("test-client/ProjectService1.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		println("Error opening OutTest.wsdl:", err.Error())
		return
	}
	defer w.Close()

	// _, err = w.WriteString("package " + wsdl.Name)
	_, err = w.WriteString("package main\n\n")
	if err != nil {
		println("Error writing header:", err.Error())
		return
	}
	_, err = w.WriteString("import \"encoding/xml\"\n\n")

	for _, xsd := range xsdMap {
		for _, element := range xsd.Elements {
			if element.ComplexType == nil {
				continue
			}
			WriteComplexType(w, element.Name, element.ComplexType, xsd)
		}
		for _, complexType := range xsd.ComplexTypes {
			WriteComplexType(w, complexType.Name, complexType, xsd)
		}
	}
}

func WriteComplexType(w io.Writer, elementName string, complexType *XsdComplexType, xsd *XsdSchema) {
	io.WriteString(w, "type ")
	io.WriteString(w, elementName)
	io.WriteString(w, " struct {\n")

	if complexType.Name == "" {
		io.WriteString(w, "\tXMLName xml.Name `xml:\"")
		io.WriteString(w, xsd.TargetNamespace)
		io.WriteString(w, " ")
		io.WriteString(w, elementName)
		io.WriteString(w, "\"`\n")
	}

	for _, element := range complexType.Sequence.Elements {
		io.WriteString(w, "\t")
		io.WriteString(w, strings.Title(element.Name))
		io.WriteString(w, " ")
		if element.MaxOccurs == "unbounded" {
			io.WriteString(w, "[]")
		}
		if element.MinOccurs == "0" {
			io.WriteString(w, "*")
		}
		io.WriteString(w, FindType(element.Type))
		io.WriteString(w, " `xml:\"")
		io.WriteString(w, xsd.TargetNamespace)
		io.WriteString(w, " ")
		io.WriteString(w, element.Name)
		io.WriteString(w, "\"`\n")
	}
	io.WriteString(w, "}\n\n")
}

func FindType(elementType string) string {
	if strings.HasPrefix(elementType, "tns:") {
		return elementType[4:]
	} else if elementType == "xs:int" {
		return "int32"
	}
	return "string"

}

func LoadWsdl(url string, wsdlMap map[string]*WsdlDefinitions, xsdMap map[string]*XsdSchema) (retval *WsdlDefinitions, err error) {
	existing := wsdlMap[url]
	if existing != nil {
		return existing, nil
	}

	println("Loading WSDL:", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	retval, err = ParseWsdl(response.Body, "http://dws-na2-int.replicon.com/", "http://na2.replicon.com/services/")
	if err != nil {
		return nil, err
	}
	wsdlMap[url] = retval

	for _, imp := range retval.Imports {
		wsdlMap[imp.Location], err = LoadWsdl(imp.Location, wsdlMap, xsdMap)
		if err != nil {
			return nil, err
		}
	}

	for _, types := range retval.Types {
		for _, schema := range types.XsdSchema {
			for _, imp := range schema.Imports {
				_, err = LoadXsd(imp.SchemaLocation, xsdMap)
			}
		}
	}

	return retval, nil
}

func LoadXsd(url string, xsdMap map[string]*XsdSchema) (retval *XsdSchema, err error) {
	existing := xsdMap[url]
	if existing != nil {
		return existing, nil
	}

	println("Loading XSD:", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	retval, err = ParseXsd(response.Body, "http://dws-na2-int.replicon.com/", "http://na2.replicon.com/services/")
	if err != nil {
		return nil, err
	}
	xsdMap[url] = retval

	for _, imp := range retval.Imports {
		_, err = LoadXsd(imp.SchemaLocation, xsdMap)
	}

	return retval, nil
}

func DebugPrintWsdl(wsdl *WsdlDefinitions) {
	println("WSDL:", wsdl.TargetNamespace, wsdl.Name)
	for _, imp := range wsdl.Imports {
		println("Import:", imp.Location, imp.Namespace)
	}
	for _, types := range wsdl.Types {
		for _, schema := range types.XsdSchema {
			println("Schema:", schema.TargetNamespace)
			for _, imp := range schema.Imports {
				println("\tImport:", imp.SchemaLocation, imp.Namespace)
			}
		}
	}
	for _, message := range wsdl.Messages {
		println("Message:", message.Name)
		for _, part := range message.Parts {
			println("\tPart:", part.Name, part.Element)
		}
	}
	for _, portType := range wsdl.PortTypes {
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
	for _, service := range wsdl.Services {
		println("Service:", service.Name)
		for _, port := range service.Ports {
			println("\tPort:", port.Name, port.Binding)
			for _, address := range port.SoapAddresses {
				println("\t\tAddress:", address.Location)
			}
		}
	}
	for _, binding := range wsdl.Bindings {
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
}

func ParseWsdl(reader io.Reader, urlSubstituteOld string, urlSubstituteNew string) (*WsdlDefinitions, error) {
	decoder := xml.NewDecoder(reader)
	obj := WsdlDefinitions{}
	err := decoder.Decode(&obj)
	if err != nil {
		return nil, err
	}

	for _, imp := range obj.Imports {
		imp.Location = strings.Replace(imp.Location, urlSubstituteOld, urlSubstituteNew, -1)
	}
	for _, types := range obj.Types {
		for _, schema := range types.XsdSchema {
			for _, imp := range schema.Imports {
				imp.SchemaLocation = strings.Replace(imp.SchemaLocation, urlSubstituteOld, urlSubstituteNew, -1)
			}
		}
	}

	return &obj, nil
}

func ParseXsd(reader io.Reader, urlSubstituteOld string, urlSubstituteNew string) (*XsdSchema, error) {
	decoder := xml.NewDecoder(reader)
	obj := XsdSchema{}
	err := decoder.Decode(&obj)
	if err != nil {
		return nil, err
	}

	for _, imp := range obj.Imports {
		imp.SchemaLocation = strings.Replace(imp.SchemaLocation, urlSubstituteOld, urlSubstituteNew, -1)
	}

	return &obj, nil
}

// 	println("WSDL:", obj.TargetNamespace, obj.Name)
// 	for _, imp := range obj.Imports {
// 		println("Import:", strings.Replace(imp.Namespace, substituteOld, substituteNew, -1), imp.Location)
// 	}
// 	for _, types := range obj.Types {
// 		for _, schema := range types.XsdSchema {
// 			println("Schema:", schema.TargetNamespace)
// 			for _, imp := range schema.Imports {
// 				println("\tImport:", strings.Replace(imp.SchemaLocation, substituteOld, substituteNew, -1), imp.Namespace)
// 			}
// 		}
// 	}
// 	for _, message := range obj.Messages {
// 		println("Message:", message.Name)
// 		for _, part := range message.Parts {
// 			println("\tPart:", part.Name, part.Element)
// 		}
// 	}
// 	for _, portType := range obj.PortTypes {
// 		println("PortType:", portType.Name)
// 		for _, operation := range portType.Operations {
// 			println("\tOperation:", operation.Name)
// 			for _, input := range operation.Inputs {
// 				println("\t\tInput:", input.Message, input.WsawAction)
// 			}
// 			for _, output := range operation.Outputs {
// 				println("\t\tOutput:", output.Message, output.WsawAction)
// 			}
// 			for _, fault := range operation.Faults {
// 				println("\t\tFault:", fault.Message, fault.WsawAction)
// 			}
// 		}
// 	}
// 	for _, service := range obj.Services {
// 		println("Service:", service.Name)
// 		for _, port := range service.Ports {
// 			println("\tPort:", port.Name, port.Binding)
// 			for _, address := range port.SoapAddresses {
// 				println("\t\tAddress:", address.Location)
// 			}
// 		}
// 	}
// 	for _, binding := range obj.Bindings {
// 		println("Binding:", binding.Name, binding.Type)
// 		for _, soapBinding := range binding.SoapBindings {
// 			println("\tSOAP Binding:", soapBinding.Transport)
// 		}
// 		for _, operation := range binding.Operations {
// 			println("\tOperation:", operation.Name)
// 			for _, soapOperation := range operation.SoapOperations {
// 				println("\t\tSOAP Operation:", soapOperation.SoapAction, soapOperation.Style)
// 			}
// 		}
// 	}

// 	/*
// 		w, err := os.OpenFile("OutTest.wsdl", os.O_CREATE|os.O_WRONLY, 0644)
// 		if err != nil {
// 			println("Error opening OutTest.wsdl:", err.Error())
// 			return
// 		}

// 		_, err = w.WriteString(xml.Header)
// 		if err != nil {
// 			println("Error writing header:", err.Error())
// 			return
// 		}

// 		encoder := xml.NewEncoder(w)
// 		err = encoder.Encode(&obj)
// 		if err != nil {
// 			println("Error encoding document:", err.Error())
// 			return
// 		}

// 		w.Close()
// 	*/
// }
