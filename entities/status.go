package entities

import "encoding/xml"

type Status struct {
	XMLName  xml.Name `json:"-"        xml:"result"`
	Database string   `json:"database" xml:"database"`
	Version  string   `json:"version"  xml:"version"`
}
