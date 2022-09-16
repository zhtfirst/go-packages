package config

import (
	"log"

	uber_config "go.uber.org/config"
)

var GlobalConfig *uber_config.YAML

func Setup(path string) {
	//指定配置文件路径
	if len(path) == 0 {
		path = "example/conf.d/conf.yaml"
	}
	var err error
	GlobalConfig, err = uber_config.NewYAML(uber_config.File(path))
	if nil != err {
		log.Panicf("read config error: %#v", err)
	}
}

func GetValue(keys ...string) uber_config.Value {
	value := GlobalConfig.Get("")
	for _, k := range keys {
		value = value.Get(k)
	}
	return value
}

func GetString(keys ...string) string {
	value := GlobalConfig.Get("")
	for _, k := range keys {
		value = value.Get(k)
	}
	return value.String()
}

func GetInt(keys ...string) (result int64) {
	value := GlobalConfig.Get("")
	for _, k := range keys {
		value = value.Get(k)
	}
	err := value.Populate(&result)
	if err != nil {
		return 0
	}
	return
}

func GetBoole(keys ...string) (result bool) {
	value := GlobalConfig.Get("")
	for _, k := range keys {
		value = value.Get(k)
	}
	err := value.Populate(&result)
	if err != nil {
		return false
	}
	return
}

func GetFloat(keys ...string) (result float64) {
	value := GlobalConfig.Get("")
	for _, k := range keys {
		value = value.Get(k)
	}
	err := value.Populate(&result)
	if err != nil {
		return 0
	}
	return
}
