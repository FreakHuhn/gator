package main

import (
	"fmt"

	"github.com/FreakHuhn/gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("FreakHuhn")
	cfg = config.Read()
	fmt.Println(cfg)
}