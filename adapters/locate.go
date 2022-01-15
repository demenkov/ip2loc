package adapters

import (
	"github.com/demenkov/ip2loc/entities"
	"github.com/ip2location/ip2location-go"
)

func Locate(ip string, db *ip2location.DB) (entities.Location, error) {
	results, err := db.Get_all(ip)
	if err != nil {
		return entities.Location{
			CodeCountry: "-",
			Country:     "-",
			County:      "-",
			City:        "-",
		}, err
	}
	return entities.Location{
		CodeCountry: results.Country_short,
		Country:     results.Country_long,
		County:      results.Region,
		City:        results.City,
	}, nil
}
