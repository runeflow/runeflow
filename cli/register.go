package main

import (
	"fmt"

	"github.com/runeflow/runeflow/api"
	"github.com/runeflow/runeflow/util"

	flag "github.com/spf13/pflag"
)

func register(a *api.API) {
	fmt.Println("Welcome to Runeflow! If you already have an account, run 'runeflow auth' instead.")
	var email, firstName, lastName string
	flag.StringVar(&email, "email", "", "Email to register with Runeflow.")
	flag.StringVar(&firstName, "first-name", "", "Your first name for your account")
	flag.StringVar(&lastName, "last-name", "", "Your last name for your account")
	flag.Parse()
	if email == "" {
		email = util.PromptString("Email: ")
	}
	if firstName == "" {
		firstName = util.PromptString("First Name: ")
	}
	if lastName == "" {
		lastName = util.PromptString("Last Name: ")
	}
	if err := a.Register(email, firstName, lastName); err != nil {
		fmt.Println("There was a problem completing your registration. The error was:")
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("Success! Next:")
	fmt.Println("1. Click the link in your email to confirm your account")
	fmt.Println("2. Run 'runeflow auth' to add this server to your account")
}
