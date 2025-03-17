package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hulkbusterks/BlogAggregator/internal/config"
	"github.com/hulkbusterks/BlogAggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config")
	}

	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database")
	}

	dbQueries := database.New(db)
	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := Commands{
		registred_commands: map[string]func(*state, Command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)

	if len(os.Args) < 2 {
		fmt.Println("inavlid command use : cli <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(s, Command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
