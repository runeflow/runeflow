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

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	cfg := config.NewConfig()
	a := api.NewAPI(cfg)
	client := client.NewClient(cfg.GetAPIKey(), cfg.GetEndpoint())
	switch os.Args[1] {
	case "run":
		run(client)
		os.Exit(0)
	case "register":
		register(a)
		os.Exit(0)
	case "auth":
		auth(a)
		os.Exit(0)
	case "-h", "--help", "help":
		usage()
		os.Exit(0)
	case "-v", "--version":
		showVersion()
		os.Exit(0)
	}
	usage()
	os.Exit(1)
}

func showVersion() {
	fmt.Println("Runeflow CLI")
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Compiled on: %s\n", buildDate)
	fmt.Printf("Git revision: %s\n", gitRevision)
	fmt.Printf("Source URL: https://github.com/runeflow/runeflow/tree/%s\n", gitRevision)
}
