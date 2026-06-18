package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Db_Url string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func(c Config) SetUser(name string) {
	c.Current_user_name = name
	homeLocation, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configFileLocation := homeLocation + "/.gatorconfig.json"
	file, err := os.Create(configFileLocation)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		panic(err)
	}
}

func Read() Config {
	homeLocation, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configFileLocation := homeLocation + "/.gatorconfig.json"
	var config Config
	file, err := os.Open(configFileLocation)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}