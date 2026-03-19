package main

import (
	"log"
	"os"
	"database/sql"

	"github.com/eric-engberg/blog-aggregator-boot.dev/internal/config"
	"github.com/eric-engberg/blog-aggregator-boot.dev/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()

	programState := &state{
		cfg: &cfg,
	}

	commands := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	programState.db = database.New(db)

	err = commands.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}
}
