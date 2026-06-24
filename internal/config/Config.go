package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Db_Url string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

type state struct {
	ConfigStruct *Config
}

type command struct {
	Name []string
}

type commands struct {
	Commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.Commands[cmd.Name[0]]; ok {
		return handler(s, cmd)
	}
	return errors.New("unknown command")
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

func handlerLogin(s *state, c command) error {
	if len(c.Name) == 0 {
		return errors.New("empty args")
	}
	if len(c.Name[0]) == 0 {
		return errors.New("no name provided")
	}
	s.ConfigStruct.SetUser(c.Name[0])
	fmt.Printf("User %s logged in\n", c.Name[0])
	return nil
}