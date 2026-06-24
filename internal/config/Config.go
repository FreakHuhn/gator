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

type State struct {
	ConfigStruct *Config
}

type Command struct {
	Name []string
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

// Führt den Befehl aus, der dem angegebenen Namen entspricht.
func (c *Commands) Run(s *State, cmd Command) error {
	if handler, ok := c.Commands[cmd.Name[0]]; ok {
		return handler(s, cmd)
	}
	return errors.New("unknown command")
}

// Registriert einen neuen Befehl mit dem angegebenen Namen und der zugehörigen Funktion.
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

// SetUser setzt den aktuellen Benutzernamen in 
// der Konfiguration und speichert die Konfiguration in einer JSON-Datei.
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

// Read liest die Konfiguration aus der JSON-Datei und gibt sie zurück.
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

// HandlerLogin ist eine Funktion, die den Login-Befehl verarbeitet.
func HandlerLogin(s *State, c Command) error {
	if len(c.Name) < 2 {
		return errors.New("username required")
	}
	username := c.Name[1]
	s.ConfigStruct.SetUser(username)
	fmt.Printf("User %s logged in\n", username)
	return nil
}