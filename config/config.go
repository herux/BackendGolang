package config

import (
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type appConfig struct {
	ApiKeyOpenWeather string       `yaml:"oweather_apikey"`
	DB                DbConfig     `yaml:"database"`
	Service           ServerConfig `yaml:"service"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Protocol string `yaml:"protocol"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type ServerConfig struct {
	Port        uint16 `yaml:"port"`
	ReadTimeout uint   `yaml:"readTimeout"`
}

var (
	k    = koanf.New(".")
	conf = appConfig{}
)

func Load(config string) error {
	err := k.Load(file.Provider(config), yaml.Parser())
	if err != nil {
		goto HandleError
	}

	if err != nil {
		goto HandleError
	}

	err = k.UnmarshalWithConf("", &conf, koanf.UnmarshalConf{Tag: "yaml"})
	if err != nil {
		goto HandleError
	}

HandleError:
	if err != nil {
		os.Exit(1)
	}

	return err
}

func DatabaseConfig() *DbConfig {
	return &conf.DB
}

func OpenWeatherAPIKey() string {
	return conf.ApiKeyOpenWeather
}

func Service() *ServerConfig {
	return &conf.Service
}
