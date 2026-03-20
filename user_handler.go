package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hyperflyguy/ThreatGator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username was supplied")
	}
	name := sql.NullString{
		String: cmd.args[0],
		Valid:  true,
	}
	// Check for name in database
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		s.c.SetUser(name.String)
		fmt.Printf("User has been set to %s", s.c.CurrentUsername)
		return nil
	}
	fmt.Println("User does not exist")
	os.Exit(1)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	//Make sure a user is supplied
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username was supplied")
	}
	// Configure a user and add them to the database
	name := sql.NullString{
		String: cmd.args[0],
		Valid:  true,
	}
	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	// Check for name in database
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Create the user
	s.db.CreateUser(context.Background(), user)
	s.c.SetUser(name.String)
	fmt.Printf("User %s was added to the Database\n", name.String)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetDatabase(context.Background())
	if err != nil {
		fmt.Println("Database was not reset")
		os.Exit(1)
	}
	os.Exit(0)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Unable to retrieve users from Database")
		os.Exit(1)
	}
	for _, user := range users {
		if user.Name.String == s.c.CurrentUsername {
			fmt.Printf("* %s (current)\n", user.Name.String)
		} else {
			fmt.Printf("* %s\n", user.Name.String)
		}
	}
	os.Exit(0)
	return nil
}
