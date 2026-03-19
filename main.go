package main

import (
	"fmt"
	"os"

	"github.com/hyperflyguy/ThreatGator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	s := state{c: &cfg}
	args := os.Args
	cmd := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmd.register("login", handlerLogin)
	if len(args) < 2 {
		fmt.Println("Invalid number of arguments (expected 2)")
		os.Exit(1)
	} else {
		cmd_run := command{
			name: args[1],
			args: args[2:],
		}
		err := cmd.run(&s, cmd_run)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
