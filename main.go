package main

import (
	"github.com/karimOCB/blog_aggregator/internal/config"
	"log"
	"fmt"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	loadedCfg, err := config.Read()
	
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	statePtr := &state{
		cfg: &loadedCfg,
	} 
	
	cmds := commands{
		registry: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments were provided")
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(statePtr, cmd)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Config Struct: %+v\n", loadedCfg)

}
