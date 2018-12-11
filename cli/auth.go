package main

import (
	"fmt"
	"os"

	"github.com/runeflow/runeflow/api"
	"github.com/runeflow/runeflow/osrelease"
	"github.com/runeflow/runeflow/util"
	flag "github.com/spf13/pflag"
)

func auth(a *api.API) {
	var email string
	flag.StringVar(&email, "email", "", "Email of the account you wish to connect this agent to.")
	flag.Parse()
	if email == "" {
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
