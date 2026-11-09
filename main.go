package main

import (
	"log"
	"os"

	"github.com/ppieniazek/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	s := &state{cfg: &cfg}
	cmds := commands{handlers: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}
	cmd := command{Name: os.Args[1], Args: os.Args[2:]}
	if err = cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}

}

type state struct {
	cfg *config.Config
}
