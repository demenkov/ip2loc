package connectors

import (
	"github.com/ip2location/ip2location-go"
	"os"
)

func LocationDb() (*ip2location.DB, error) {
	var db *ip2location.DB
	var err error
	db, err = ip2location.OpenDB(os.Getenv("IPDB"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
