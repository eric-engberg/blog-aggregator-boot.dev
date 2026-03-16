package main

import (
	"fmt"
	"log"

	"github.com/eric-engberg/blog-aggregator-boot.dev/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = cfg.SetCurrentUserName("eric")
	if err != nil {
		log.Fatalf("error setting current user name: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)
}
