package helpers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func FormatResponse(w http.ResponseWriter, format string, resp interface{}) {
	header := w.Header()
	header.Set("Content-Type", fmt.Sprintf("application/%s; charset=utf-8", format))
	if format == "xml" {
		xml.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
