package main

import (
	"fmt"

	"github.com/FreakHuhn/gator/internal/config"

func main() {
	config.Read()
	config.SetUser()
	fmt.Println(config.Read())
}