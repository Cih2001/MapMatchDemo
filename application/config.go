package application

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/labstack/echo/middleware"
)

var (
	Config      = loadConfig()
	configPath  = "config.json"
	defaultConf = Cgf{
		MongoURL:      "localhost:27017",
		LogPath:       "common.log",
		Debug:         false,
		DatabaseName:  "GoTruck",
		ListenAddress: "localhost:7080",
		Views:         "views/",
	}
)

type Cgf struct {
	MongoURL      string `json:"mongo_url"`
	LogPath       string `json:"log_path"`
	Debug         bool   `json:"debug"`
	DatabaseName  string `json:"database_name"`
	ListenAddress string `json:"listen_address"`
	Views         string `json:"views"`
}

func (c *Cgf) requiredColumnExists() bool {
	allExists := true
	cType := reflect.TypeOf(c)
	cVal := reflect.ValueOf(c)
	if cVal.Kind() == reflect.Ptr {
		cVal = cVal.Elem()
		cType = cType.Elem()
	}
	for i := 0; i < cVal.NumField(); i++ {
		field := cType.Field(i)
		val, ok := field.Tag.Lookup("required")
		if !ok || val == "false" {
			continue
		}
		value := cVal.Field(i)
		switch value.Type().Kind() {
		case reflect.Bool:
			continue
		case reflect.String:
			if value.String() == "" {
				allExists = false
			}
		case reflect.Int:
			if value.Int() == 0 {
				allExists = false
			}
		default:
			allExists = false
		}
	}
	return allExists
}

func loadConfig() *Cgf {
	cgf := &defaultConf
	file, err := ioutil.ReadFile(configPath)
	if err == nil {
		json.Unmarshal(file, cgf)
	}
	if !cgf.requiredColumnExists() {
		panic("Loading Config required faild")
	}
	return cgf
}

func loadLoggerConfig(config *Cgf) middleware.LoggerConfig {
	logConf := middleware.DefaultLoggerConfig
	f, err := os.OpenFile(config.LogPath, os.O_CREATE|os.O_RDWR, os.FileMode(0666))
	if err != nil {
		panic(err.Error())
	}

	logConf.Output = f
	return logConf
}
