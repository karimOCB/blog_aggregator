package main

import (
	"github.com/karimOCB/blog_aggregator/internal/config"
	"log"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = cfg.SetUser("Karim")

	if err != nil {
		log.Fatalf("error setting user in config: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	fmt.Printf("Config Struct: %+v\n", cfg)

}
