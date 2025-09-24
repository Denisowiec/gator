package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Denisowiec/gator/internal/config"
	"github.com/Denisowiec/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error occured: %v", err)
	}

	db, err := sql.Open("postgres", cfg.Dburl)
	if err != nil {
		fmt.Printf("Could not connect to the database")
		os.Exit(1)
	}
	dbQueries := database.New(db)

	programState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	programCommands := commands{}

	programCommands.register("login", handlerLogin)
	programCommands.register("register", handlerRegister)
	programCommands.register("reset", handlerReset)
	programCommands.register("users", handlerUsers)

	if len(os.Args) < 2 {
		fmt.Println("no command given")
		os.Exit(1)
	} else {
		cmdToRun := command{
			name: os.Args[1],
			args: []string{},
		}
		if len(os.Args) > 2 {
			cmdToRun.args = os.Args[2:]
		}
		if err := programCommands.run(&programState, cmdToRun); err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}
}
