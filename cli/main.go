package main

import (
	"fmt"
	"os"

	"github.com/runeflow/runeflow/api"
	"github.com/runeflow/runeflow/client"
	"github.com/runeflow/runeflow/config"
)

var version string
var buildDate string
var gitRevision string

const (
	// ExitStatusOK indicates successful execution
	ExitStatusOK = 0

	// ExitStatusInvalidCommand indicates an unrecognized usage
	ExitStatusInvalidCommand = 1

	// ExitStatusNoAgentID indicates that Runeflow tried to run, but an Agent ID
	// was not configured
	ExitStatusNoAgentID = 2

	// ExitStatusAPIError indicates that there was an API error when running the
	// command
	ExitStatusAPIError = 3
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(ExitStatusInvalidCommand)
	}
	cfg := config.NewConfig()
	a := api.NewAPI(cfg)
	switch os.Args[1] {
	case "run":
		agentID, err := cfg.GetAgentID()
		if err != nil {
			fmt.Printf("error getting agent id: %v\n", err)
			os.Exit(ExitStatusNoAgentID)
		}
		run(client.NewClient(agentID, cfg.GetEndpoint()))
		os.Exit(ExitStatusOK)
	case "register":
		register(a)
		os.Exit(ExitStatusOK)
	case "auth":
		auth(a)
		os.Exit(ExitStatusOK)
	case "-h", "--help", "help":
		usage()
		os.Exit(ExitStatusOK)
	case "-v", "--version":
		showVersion()
		os.Exit(ExitStatusOK)
	}
	usage()
	os.Exit(ExitStatusInvalidCommand)
}

func showVersion() {
	fmt.Println("Runeflow CLI")
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Compiled on: %s\n", buildDate)
	fmt.Printf("Git revision: %s\n", gitRevision)
	fmt.Printf("Source URL: https://github.com/runeflow/runeflow/tree/%s\n", gitRevision)
}
