package response

import (
	"github.com/demenkov/ip2loc/entities"
	"github.com/demenkov/ip2loc/helpers"
	"net/http"
)

func Error(w http.ResponseWriter, format string, errorCode int) {
	w.WriteHeader(errorCode)
	e := &entities.Error{
		Code:    errorCode,
		Message: http.StatusText(errorCode),
	}
	helpers.FormatResponse(w, format, e)
}
