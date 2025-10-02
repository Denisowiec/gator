package main

import "fmt"

func handlerHelp(s *state, cmd command) error {
	helpList := make(map[string]string)

	helpList["login"] = "login <username>: log in as <username>"
	helpList["register"] = "register <username>: register a user"
	helpList["reset"] = "deletes all users from database"
	helpList["users"] = "lists all registered users"
	helpList["agg"] = "agg <time>: starts the aggregation process. <time> allows you to select time interval between aggregations"
	helpList["addfeed"] = "addfeed <url>: adds a feed at <url> to the database and follows it for current user"
	helpList["feeds"] = "lists all feeds in the database"
	helpList["follow"] = "follow <url>: follow a feed as the current user. The feed must first be added to the database"
	helpList["following"] = "lists all feeds followed by the current user"
	helpList["unfollow"] = "unfollow <url>: unfollow a feed"
	helpList["browse"] = "browse <limit>: lists the recent posts from feeds followed by the current user, limitted by <limit> (defaults to 2)"
	helpList["help"] = "help <cmd>: Display help for command <cmd>"

	fmt.Println("Usage: gator <cmd> <arguments>")
	if len(cmd.args) < 1 {
		fmt.Println("Available commands:")
		for key, val := range helpList {
			fmt.Printf("%v - %v\n", key, val)
		}
	} else {
		if t, ok := helpList[cmd.args[0]]; ok {
			fmt.Printf("%v - %v\n", cmd.args[0], t)
		} else {
			fmt.Println("unknown command")
		}
	}

	return nil
}
