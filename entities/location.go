package entities

import "encoding/xml"

type Location struct {
	XMLName     xml.Name `json:"-"            xml:"result"`
	CodeCountry string   `json:"code_country" xml:"code_country"`
	Country     string   `json:"country"      xml:"country"`
	County      string   `json:"county"       xml:"county"`
	City        string   `json:"city"         xml:"city"`
}
