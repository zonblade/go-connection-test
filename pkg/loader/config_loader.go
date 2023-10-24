package loader

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	DB       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"dbName"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func loadConfig() ([]Data, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile("db.config.json")
	if err != nil {
		return nil, err
	}

	// Create a slice of Person to store the unmarshaled data
	var people []Data

	// Unmarshal the JSON data into the slice
	if err := json.Unmarshal(data, &people); err != nil {
		return nil, err
	}
	return people, nil
}

func GetDbConfig(dbName string) []Data {
	var filteredData []Data

	dataList, err := loadConfig()

	if err != nil {
		panic("no config available")
	}

	for _, data := range dataList {
		if data.DB == dbName {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}
