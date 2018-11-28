package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/runeflow/runeflow/api"
)

// RegisterJSON is a registration message
type RegisterJSON struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func register(a *api.API) {
	fmt.Println("Welcome to Runeflow! If you already have an account, run 'runeflow auth' instead.")
	var msg RegisterJSON
	reader := bufio.NewReader(os.Stdin)
	promptString(reader, "Email: ", &msg.Email)
	promptString(reader, "First Name: ", &msg.FirstName)
	promptString(reader, "Last Name: ", &msg.LastName)
	if err := a.Register(msg); err != nil {
		fmt.Println("There was a problem completing your registration. The error was:")
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("Success! Next:")
	fmt.Println("1. Click the link in your email to confirm your account")
	fmt.Println("2. Run 'runeflow auth' to add this server to your account")
}
