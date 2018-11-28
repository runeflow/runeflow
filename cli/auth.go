package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/runeflow/runeflow/api"
)

func auth(a *api.API) {
	if len(os.Args) == 4 && os.Args[2] == "--email" {
		authEmail(a, os.Args[3])
		return
	}
	reader := bufio.NewReader(os.Stdin)
	var email string
	promptString(reader, "Email: ", &email)
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
