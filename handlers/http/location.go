package http

import (
	"github.com/demenkov/ip2loc/adapters"
	"github.com/demenkov/ip2loc/connectors"
	"github.com/demenkov/ip2loc/helpers"
	"github.com/demenkov/ip2loc/response"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

func Location(w http.ResponseWriter, r *http.Request) {
	db, _ := connectors.LocationDb()
	vars := mux.Vars(r)
	ip := vars["ip"]
	format := vars["format"]

	if net.ParseIP(ip) == nil {
		response.Error(w, format, http.StatusBadRequest)
		return
	}
	loc, err := adapters.Locate(ip, db)
	if err != nil {
		response.Error(w, format, http.StatusNotFound)
		return
	}
	helpers.FormatResponse(w, format, loc)
}
