package entities

import "encoding/xml"

type Error struct {
	XMLName xml.Name `json:"-"       xml:"result"`
	Code    int      `json:"code"    xml:"code"`
	Message string   `json:"message" xml:"message"`
}
