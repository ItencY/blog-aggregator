package main

import (
	"log"
	"os"

	"github.com/itency/blog_aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := state{
		cfg: &cfg,
	}

	cmds := commands{
		commands: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
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
