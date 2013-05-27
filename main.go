package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type WsdlDefinitions struct {
	Name            string           `xml:"name,attr"`
	TargetNamespace string           `xml:"targetNamespace,attr"`
	Imports         []*WsdlImport    `xml:"http://schemas.xmlsoap.org/wsdl/ import"`
	Types           []*WsdlTypes     `xml:"http://schemas.xmlsoap.org/wsdl/ types"`
	Messages        []*WsdlMessage   `xml:"http://schemas.xmlsoap.org/wsdl/ message"`
	PortTypes       []*WsdlPortTypes `xml:"http://schemas.xmlsoap.org/wsdl/ portType"`
	Services        []*WsdlService   `xml:"http://schemas.xmlsoap.org/wsdl/ service"`
	Bindings        []*WsdlBinding   `xml:"http://schemas.xmlsoap.org/wsdl/ binding"`
}

type WsdlBinding struct {
	Name         string           `xml:"name,attr"`
	Type         string           `xml:"type,attr"`
	Operations   []*WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
	SoapBindings []*SoapBinding   `xml:"http://schemas.xmlsoap.org/wsdl/soap/ binding"`
}

type SoapBinding struct {
	Transport string `xml:"transport,attr"`
}

type WsdlTypes struct {
	XsdSchema []*XsdSchema `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

type WsdlImport struct {
	Namespace string `xml:"namespace,attr"`
	Location  string `xml:"location,attr"`
}

type WsdlMessage struct {
	Name  string             `xml:"name,attr"`
	Parts []*WsdlMessagePart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

type WsdlMessagePart struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

type WsdlPortTypes struct {
	Name       string           `xml:"name,attr"`
	Operations []*WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

type WsdlOperation struct {
	Name           string                 `xml:"name,attr"`
	Inputs         []*WsdlOperationInput  `xml:"http://schemas.xmlsoap.org/wsdl/ input"`
	Outputs        []*WsdlOperationOutput `xml:"http://schemas.xmlsoap.org/wsdl/ output"`
	Faults         []*WsdlOperationFault  `xml:"http://schemas.xmlsoap.org/wsdl/ fault"`
	SoapOperations []*SoapOperation       `xml:"http://schemas.xmlsoap.org/wsdl/soap/ operation"`
}

type WsdlOperationInput struct {
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlOperationOutput struct {
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlOperationFault struct {
	Name       string `xml:"name,attr"`
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlService struct {
	Name  string      `xml:"name,attr"`
	Ports []*WsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}

type WsdlPort struct {
	Name          string         `xml:"name,attr"`
	Binding       string         `xml:"binding,attr"`
	SoapAddresses []*SoapAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
}

type SoapAddress struct {
	Location string `xml:"location,attr"`
}

type SoapOperation struct {
	SoapAction string `xml:"soapAction,attr"`
	Style      string `xml:"style,attr"`
}

type SoapBody struct {
	Use string `xml:"use,attr"`
}

type XsdSchema struct {
	TargetNamespace    string            `xml:"targetNamespace,attr"`
	ElementFormDefault string            `xml:"elementFormDefault,attr"`
	Imports            []*XsdImport      `xml:"http://www.w3.org/2001/XMLSchema import"`
	Elements           []*XsdElement     `xml:"http://www.w3.org/2001/XMLSchema element"`
	ComplexTypes       []*XsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type XsdImport struct {
	SchemaLocation string `xml:"schemaLocation,attr"`
	Namespace      string `xml:"namespace,attr"`
}

type XsdElement struct {
	Name        string          `xml:"name,attr"`
	Nillable    bool            `xml:"nillable,attr"`
	Type        string          `xml:"type,attr"`
	MinOccurs   string          `xml:"minOccurs,attr"`
	MaxOccurs   string          `xml:"maxOccurs,attr"`
	ComplexType *XsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	SimpleType  *XsdSimpleType  `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

type XsdComplexType struct {
	Name     string       `xml:"name,attr"`
	Sequence *XsdSequence `xml:"http://www.w3.org/2001/XMLSchema sequence"`
}

type XsdSimpleType struct {
	Name     string          `xml:"name,attr"`
	Sequence *XsdRestriction `xml:"http://www.w3.org/2001/XMLSchema restriction"`
}

type XsdSequence struct {
	Elements []*XsdElement `xml:"http://www.w3.org/2001/XMLSchema element"`
}

type XsdRestriction struct {
	Base         string           `xml:"base,attr"`
	Pattern      *XsdPattern      `xml:"http://www.w3.org/2001/XMLSchema pattern"`
	MinInclusive *XsdMinInclusive `xml:"http://www.w3.org/2001/XMLSchema minInclusive"`
	MaxInclusive *XsdMaxInclusive `xml:"http://www.w3.org/2001/XMLSchema maxInclusive"`
}

type XsdPattern struct {
	Value string `xml:"value,attr"`
}

type XsdMinInclusive struct {
	Value string `xml:"value,attr"`
}

type XsdMaxInclusive struct {
	Value string `xml:"value,attr"`
}

func main() {
	flag.Parse()
	wsdlRoot := flag.Arg(0)

	wsdlMap := make(map[string]*WsdlDefinitions)
	xsdMap := make(map[string]*XsdSchema)

	// "http://na2.replicon.com/services/ProjectService1.svc?wsdl"
	_, err := LoadWsdl(wsdlRoot, nil, wsdlMap, xsdMap)
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

func GetAbsoluteUrl(rawUrl string, relativeRawUrl *string) (*url.URL, error) {
	targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	if relativeRawUrl != nil {
		relativeUrl, err := url.Parse(*relativeRawUrl)
		if err != nil {
			return nil, err
		}
		targetUrl = relativeUrl.ResolveReference(targetUrl)
	}

	return targetUrl, nil
}

func LoadUrl(absoluteUrl url.URL) (io.Reader, error) {
	if absoluteUrl.Scheme == "http" || absoluteUrl.Scheme == "https" {
		response, err := http.Get(absoluteUrl.String())
		if err != nil {
			return nil, err
		}
		return response.Body, nil
	} else if absoluteUrl.Scheme == "file" {
		if absoluteUrl.Host != "" && absoluteUrl.Host != "localhost" {
			return nil, errors.New("file:// URL has unknown host " + absoluteUrl.Host)
		}
		path := absoluteUrl.String()[5:]
		return os.Open(path)
	}
	return nil, errors.New("unable to understand url scheme:" + absoluteUrl.Scheme)
}

func LoadWsdl(url string, relativeUrl *string, wsdlMap map[string]*WsdlDefinitions, xsdMap map[string]*XsdSchema) (retval *WsdlDefinitions, err error) {
	absoluteUrl, err := GetAbsoluteUrl(url, relativeUrl)
	if err != nil {
		return nil, err
	}

	existing := wsdlMap[absoluteUrl.String()]
	if existing != nil {
		return existing, nil
	}

	println("Loading WSDL:", absoluteUrl.String())
	response, err := LoadUrl(*absoluteUrl)
	if err != nil {
		return nil, err
	}

	retval, err = ParseWsdl(response, "http://dws-na2-int.replicon.com/", "http://na2.replicon.com/services/")
	if err != nil {
		return nil, err
	}
	wsdlMap[absoluteUrl.String()] = retval

	absoluteUrlString := absoluteUrl.String()
	for _, imp := range retval.Imports {
		wsdlMap[imp.Location], err = LoadWsdl(imp.Location, &absoluteUrlString, wsdlMap, xsdMap)
		if err != nil {
			return nil, err
		}
	}

	for _, types := range retval.Types {
		for _, schema := range types.XsdSchema {
			for _, imp := range schema.Imports {
				_, err = LoadXsd(imp.SchemaLocation, &absoluteUrlString, xsdMap)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return retval, nil
}

func LoadXsd(url string, relativeUrl *string, xsdMap map[string]*XsdSchema) (retval *XsdSchema, err error) {
	absoluteUrl, err := GetAbsoluteUrl(url, relativeUrl)
	if err != nil {
		return nil, err
	}

	existing := xsdMap[absoluteUrl.String()]
	if existing != nil {
		return existing, nil
	}

	println("Loading XSD:", absoluteUrl.String())
	response, err := LoadUrl(*absoluteUrl)
	if err != nil {
		return nil, err
	}

	retval, err = ParseXsd(response, "http://dws-na2-int.replicon.com/", "http://na2.replicon.com/services/")
	if err != nil {
		return nil, err
	}
	xsdMap[absoluteUrl.String()] = retval

	absoluteUrlString := absoluteUrl.String()
	for _, imp := range retval.Imports {
		_, err = LoadXsd(imp.SchemaLocation, &absoluteUrlString, xsdMap)
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
