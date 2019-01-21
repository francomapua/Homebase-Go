package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ConfigJSON : Class
type ConfigJSON struct {
	IntField    int    `json:"intField"`
	StringField string `json:"stringField"`
	ObjectField struct {
		One   int     `json:"one"`
		Two   string  `json:"two"`
		Three float64 `json:"three"`
	} `json:"objectField"`
}

// NewFromFile : Constructor
func NewFromFile(fpath string) (*ConfigJSON, error) {
	// Instantiate new ConfigJSON
	config := new(ConfigJSON)

	// Read File
	byteArr, err := readFileAsByteArray(fpath)
	if err != nil {
		return nil, err
	}

	// Parse JSON into struct
	json.Unmarshal(byteArr, &config)
	return config, nil
}

func readFileAsByteArray(filePath string) ([]byte, error) {
	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()
	byteArr, err := ioutil.ReadAll(jsonFile)
	return byteArr, err
}
