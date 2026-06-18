package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/itency/blog_aggregator/internal/config"
	"github.com/itency/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		commands: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	name := os.Args[1]
	args := os.Args[2:]

	err = cmds.run(&programState, command{name: name, args: args})
	if err != nil {
		log.Fatal(err)
	}
}
