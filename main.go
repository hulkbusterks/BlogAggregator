package main

import (
	"fmt"
	"log"

	"github.com/hulkbusterks/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config")
	}

	cfg.SetUser("hulky")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config")
	}
	fmt.Println(cfg)
}
