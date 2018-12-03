package main

import (
	"fmt"
	"os"

	"github.com/runeflow/runeflow/api"
	"github.com/runeflow/runeflow/util"
)

func auth(a *api.API) {
	var email string
	if len(os.Args) == 4 && os.Args[2] == "--email" {
		email = os.Args[3]
	} else {
		email = util.PromptString("Email: ")
	}
	if err := a.Authorize(email); err != nil {
		fmt.Println("There was a problem authorizing this server. The error was:")
		fmt.Printf("%v\n", err)
		os.Exit(ExitStatusAPIError)
	}
}
