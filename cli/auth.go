package main

import (
	"fmt"
	"os"

	"github.com/runeflow/runeflow/api"
	"github.com/runeflow/runeflow/util"
)

func auth(a *api.API) {
	if len(os.Args) == 4 && os.Args[2] == "--email" {
		authEmail(a, os.Args[3])
		return
	}
	var email string
	util.PromptString(&email, "Email: ")
	authEmail(a, email)
}

func authEmail(a *api.API, email string) {
	if err := a.Authorize(email); err != nil {
		fmt.Println("There was a problem authorizing this server. The error was:")
		fmt.Printf("%v\n", err)
		os.Exit(1)
		return
	}
}
