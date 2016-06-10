package utils;

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type Configuration struct {
	AWS          string
	JwtKey       string
	IgnoreString string
}

type ConfigLoader interface {
	Load() (string, error)
}

type JsonFileSource struct {
	fileName string
}

type JsonSource struct {
	JSON string
}

func (s JsonFileSource) Load() Configuration {
	file, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}


//
//func loadConfig(config *Configuration) (file string) {
//	data, err := ioutil.ReadFile(file)
//	if err != nil {
//		panic("Config error: Readfile error: " + err.Error())
//	}
//
//	res := make(Config)
//	if err = json.Unmarshal(data, &res); err != nil {
//		panic("Config error: Parse error: " + err.Error())
//	}
//
//	return res
//}