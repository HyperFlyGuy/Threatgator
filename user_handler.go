package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username was supplied")
	}
	s.c.SetUser(cmd.args[0])
	fmt.Printf("User has been set to %s", s.c.CurrentUsername)
	return nil
}

func handlerRegister() {

}

func handlerUsers() {

}
