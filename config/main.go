package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf Config

type Config struct {
	Port int    `json:"port"`
	Key  string `json:"key"`
	DB         struct {
		Host     string `json:"host"`
		Port     uint   `json:"port"`
		User     string `json:"usr"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Prefix   string `json:"prefix"`
	} `json:"db"`
}

func IniConfig() error {
	source, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &Conf)
	if err != nil {
		panic(err)
	}
	return nil
}
