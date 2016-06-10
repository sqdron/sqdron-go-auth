package configuration
//
//import (
//	"io/ioutil"
//	"log"
//	"encoding/json"
//)
//
//type Configuration struct {
//	AWS          awsConfig
//	Facebook facebookConfig
//	JwtKey       string `json:"jwt_key"`
//	IgnoreString string `json:"-"`
//}
//
//type awsConfig struct {
//	AccessKey    string `json:"key"`
//	AccessSecret string `json:"secret"`
//	Bucket       string `json:"bucket"`
//}
//
//type facebookConfig struct {
//	ClientID    string `json:"key"`
//	ClientSecret string `json:"secret"`
//	Bucket       string `json:"bucket"`
//}
//
//func LoadConfig(path string) Configuration {
//	file, err := ioutil.ReadFile(path)
//	if err != nil {
//		log.Fatal("Config File Missing. ", err)
//	}
//
//	var config Configuration
//	err = json.Unmarshal(file, &config)
//	if err != nil {
//		log.Fatal("Config Parse Error: ", err)
//	}
//
//	return config
//}