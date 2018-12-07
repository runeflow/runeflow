package main

import (
	"fmt"
	"os"

	"github.com/runeflow/runeflow/api"
	"github.com/runeflow/runeflow/osrelease"
	"github.com/runeflow/runeflow/util"
)

func auth(a *api.API) {
	var email string
	if len(os.Args) == 4 && os.Args[2] == "--email" {
		email = os.Args[3]
	} else {
		email = util.PromptString("Email: ")
	}
	if err := a.Authorize(email, getHostname(), getOSName()); err != nil {
		fmt.Println("There was a problem authorizing this server. The error was:")
		fmt.Printf("%v\n", err)
		os.Exit(ExitStatusAPIError)
	}
}

func getHostname() string {
	hn, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hn
}

func getOSName() string {
	n, err := osrelease.ReadField("PRETTY_NAME")
	if err != nil {
		return "unknown"
	}
	return n
}
