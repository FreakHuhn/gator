package main

import (
	"fmt"
	"os"

	"github.com/FreakHuhn/gator/internal/config"
)

func main() {
	var state config.State
	cfg := config.Read()
	state.ConfigStruct = &cfg
	commands := config.Commands{
		Commands: make(map[string]func(*config.State, config.Command) error),
	}
	commands.Register("login", config.HandlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments provided.")
		os.Exit(1)
	}
	cmd := config.Command{
		Name: os.Args[1:],
	}
	err := commands.Run(&state, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}