package app

import (
	"encoding/json"
	"io/ioutil"
)

type Config map[string]interface{}

func (a *Application) FromJsonFile(file string) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Config error: Readfile error: " + err.Error())
	}

	res := make(Config)
	if err = json.Unmarshal(data, &res); err != nil {
		panic("Config error: Parse error: " + err.Error())
	}

	return
}

type AppConfig struct{}

func (c *AppConfig) Read(filename string) {

}

func New() *AppConfig {
	return &AppConfig{}
}

type ConfigReader interface {
	Read(filename string)
}


//
//// GetMap return map config value
//func (conf Config) GetMap(path string) map[string]interface{} {
//
//	result := conf.Get(path)
//	if result == nil {
//		return map[string]interface{}{}
//	}
//
//	// Если строка, то это результат
//	switch val := result.(type) {
//	case map[string]interface{}: return val
//
//	default:
//		return map[string]interface{}{}
//	}
//}
//
//// GetString return string config value
//func (conf Config) GetString(path string) string {
//
//	result := conf.Get(path)
//	if result == nil {
//		return ""
//	}
//
//	// Если строка, то это результат
//	switch val := result.(type) {
//	case string: return val
//
//	default:
//		return ""
//	}
//}
//
//// GetArray return array config value
//func (conf Config) GetArray(path string) []interface{} {
//
//	result := conf.Get(path)
//	if result == nil {
//		return []interface{}{}
//	}
//
//	// Если строка, то это результат
//	switch val := result.(type) {
//	case []interface{}: return val
//
//	default:
//		return []interface{}{}
//	}
//}
//
//// GetBool return bool config value
//func (conf Config) GetBool(path string) bool {
//
//	result := conf.Get(path)
//	if result == nil {
//		return false
//	}
//
//	// Если строка, то это результат
//	switch val := result.(type) {
//	case bool: return val
//
//	default:
//		return false
//	}
//}
//
//// GetFloat64 return float64 config value
//func (conf Config) GetFloat64(path string) float64 {
//
//	result := conf.Get(path)
//	if result == nil {
//		return 0
//	}
//
//	// Если строка, то это результат
//	switch val := result.(type) {
//	case float64: return val
//	case int: return float64(val)
//	case int64: return float64(val)
//	case json.Number:
//		if res, err := strconv.ParseFloat(string(val), 64); err != nil {
//			return 0
//		} else {
//			return res
//		}
//
//	default:
//		return 0
//	}
//}
//
//// GetNumber return string config value
//func (conf Config) GetNumber(path string) string {
//
//	result := conf.Get(path)
//	if result == nil {
//		return ""
//	}
//
//	// Если строка, то это результат
//	switch val := result.(type) {
//	case string: return val
//	case json.Number: return string(val)
//
//	default:
//		return ""
//	}
//}
//
//// GetInt64 return int64 config value
//func (conf Config) GetInt64(path string) int64 {
//
//	result := conf.GetFloat64(path)
//	return int64(result)
//}
//
//// GetInt return int config value
//func (conf Config) GetInt(path string) int {
//
//	result := conf.GetFloat64(path)
//	return int(result)
//}
//
//// Get return config value by dotted path in json tree. Path should be like this "root.option.item"
//func (conf Config) Get(path string) interface{} {
//	items := strings.Split(path, ".")
//
//	idx := 0
//	value := map[string]interface{}(conf)
//
//	// Перебор до предпоследнего элемента
//	for idx < len(items) - 1 {
//		tmp, ok := value[items[idx]]
//		if !ok {
//			return ""
//		}
//
//		value = tmp.(map[string]interface{})
//		idx++
//	}
//
//	// Последний элемент
//	return value[items[idx]]
//}
//
//func get_cinfig(path string) string {
//	data, err := ioutil.ReadFile("config.json");
//}


func loadConfig(file string) Config {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Config error: Readfile error: " + err.Error())
	}

	res := make(Config)
	if err = json.Unmarshal(data, &res); err != nil {
		panic("Config error: Parse error: " + err.Error())
	}

	return res
}
//
//func configLoader(pageGetterFunc PageGetter) {
//	// ...
//	content := pageGetterFunc(BASE_URL)
//	// ...
//}