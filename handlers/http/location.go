package http

import (
	"github.com/demenkov/ip2loc/adapters"
	"github.com/demenkov/ip2loc/helpers"
	"github.com/demenkov/ip2loc/response"
	"github.com/gorilla/mux"
	"github.com/ip2location/ip2location-go"
	"net"
	"net/http"
)

type LocationHandler struct {
	Db *ip2location.DB
}

func (locHandler *LocationHandler) Location(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]
	format := vars["format"]

	if net.ParseIP(ip) == nil {
		response.Error(w, format, http.StatusBadRequest)
		return
	}
	loc, err := adapters.Locate(ip, locHandler.Db)
	if err != nil {
		response.Error(w, format, http.StatusNotFound)
		return
	}
	helpers.FormatResponse(w, format, loc)
}
