package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/hyperflyguy/ThreatGator/internal/config"
	"github.com/hyperflyguy/ThreatGator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// Reading the config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	//Database connection
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Println("Database connection failed")
		os.Exit(1)
	}
	dbQueries := database.New(db)
	//Setting the state
	s := state{c: &cfg, db: dbQueries}
	args := os.Args
	cmd := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	// Register the login command
	cmd.register("login", handlerLogin)
	cmd.register("register", handlerRegister)
	cmd.register("reset", handlerReset)
	cmd.register("users", handlerUsers)
	cmd.register("agg", handlerAgg)
	cmd.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmd.register("feeds", handlerListFeeds)
	cmd.register("follow", middlewareLoggedIn(handlerFollowFeed))
	cmd.register("following", middlewareLoggedIn(handlerFollowingFeed))
	cmd.register("unfollow", middlewareLoggedIn(handlerUnfollowingFeed))
	cmd.register("browse", middlewareLoggedIn(handlerBrowse))
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
