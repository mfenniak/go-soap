package main

import "encoding/xml"

type GetAllFilterDefinitions struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetAllFilterDefinitions"`
}

type GetAllFilterDefinitionsResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetAllFilterDefinitionsResponse"`
}

type GetAllColumns struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetAllColumns"`
}

type GetAllColumnsResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetAllColumnsResponse"`
}

type GetRowCounts struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetRowCounts"`
}

type GetRowCountsResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetRowCountsResponse"`
}

type GetTotals struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetTotals"`
}

type GetTotalsResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetTotalsResponse"`
}

type GetData struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetData"`
}

type GetDataResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetDataResponse"`
}

type GetAllGroupDefinitions struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetAllGroupDefinitions"`
}

type GetAllGroupDefinitionsResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetAllGroupDefinitionsResponse"`
}

type GetDataByGroup struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetDataByGroup"`
}

type GetDataByGroupResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetDataByGroupResponse"`
}

type GetRowCountByGroup struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetRowCountByGroup"`
}

type GetRowCountByGroupResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetRowCountByGroupResponse"`
}
