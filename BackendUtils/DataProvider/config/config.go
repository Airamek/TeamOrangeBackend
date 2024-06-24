package config

import (
	"fmt"
	"log"
)
import "os"
import "gopkg.in/yaml.v3"

func GetConfig(name string, out interface{}) error {
	var err error

	_, err = os.Stat(fmt.Sprintf("%s.yaml", name))

	if err != nil {
		return err
	}

	var data []byte
	data, err = os.ReadFile(fmt.Sprintf("%s.yaml", name))
	if err != nil {
		log.Fatalf("File exists, but it couldn't be opened")
	}
	log.Printf("Loading config file(%s.yaml)...\n", name)
	yaml.Unmarshal(data, out)
	return nil
}

func CreateDefaultConf(name string, data interface{}) interface{} {
	var err error
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s.yaml doesn't exist, creating it now\n", name))

	//write out structure
	var file, _ = os.Create(fmt.Sprintf("%s.yaml", name))
	var yamlData []byte
	yamlData, err = yaml.Marshal(data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	_, err = file.Write(yamlData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return data
}
