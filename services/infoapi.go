package services

import (
	"encoding/json"
	"net/http"

	"../models"
)

const (
	enterprise  = "iDev Soluções"
	about       = "API developed to read a JSON file showing (avarage, trend and usage trend) of servers"
	projectLink = "github.com/rafaelsanzio/iDev"
	creator     = "Rafael Sanzio"
)

// GetInfoAPI to return information about API
func GetInfoAPI(w http.ResponseWriter, r *http.Request) {
	getInfo := models.GetInfo{
		Enterprise:  enterprise,
		About:       about,
		ProjectLink: projectLink,
		Creator:     creator,
	}

	json.NewEncoder(w).Encode(getInfo)
}
