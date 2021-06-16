package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"log"
)

const (
	configBaseDir = "../../../../resources/config/"
)

var (
	properties Properties
)

func init() {
	fileNames := getConfigFileNames()

	for _, fileName := range fileNames {
		filePath := configBaseDir + fileName
		file, err := Asset(filePath)
		if err != nil {
			log.Printf("[CONFIG-ERROR] Fail to load configuration: %v %v", fileName, err)
			continue
		}

		err = yaml.Unmarshal(file, &properties)
		if err != nil {
			log.Printf("[CONFIG-ERROR] Fail to parse configuration: %v %v", fileName, err)
			continue
		}

		log.Printf("[CONFIG-INFO] Success on load configuration: %v:", fileName)
	}
	printConfig()
}

func printConfig() {
	bytes, _ := json.MarshalIndent(properties, "", "    ")
	log.Println("Configuration: " + string(bytes))
}

func getConfigFileNames() []string {
	fileNames := make([]string, 0)
	fileNames = append(fileNames, "app.yaml")

	return fileNames
}

func GetProps() *Properties {
	return &properties
}
