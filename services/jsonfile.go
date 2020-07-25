package services

import (
	"encoding/json"
	"io/ioutil"

	"../models"
)

//readJSONFile to read information in JSON file and Unmarshall into model
func readJSONFile(filename string) ([]models.JSONFile, error) {
	JSONData := []models.JSONFile{}
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(dat), &JSONData)
	if err != nil {
		return nil, err
	}

	return JSONData, nil
}
