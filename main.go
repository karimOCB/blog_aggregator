package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/karimOCB/blog_aggregator/internal/config"
	"github.com/karimOCB/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	loadedCfg, err := config.Read()

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", loadedCfg.DbUrl)

	if err != nil {
		log.Fatalf("error opening the database: %v", err)
	}

	dbQueries := database.New(db)

	statePtr := &state{
		cfg: &loadedCfg,
		db:  dbQueries,
	}

	defer db.Close()

	cmds := commands{
		registry: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
