package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/runeflow/runeflow/api"
	"github.com/runeflow/runeflow/client"
	"github.com/runeflow/runeflow/config"
)

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
	case "-h":
	case "--help":
	case "help":
		usage()
		os.Exit(0)
	}
	usage()
	os.Exit(1)
}

func promptString(r *bufio.Reader, prompt string, dest *string) {
	for *dest == "" {
		fmt.Print(prompt)
		if text, err := r.ReadString('\n'); err == nil {
			*dest = strings.TrimSpace(text)
		}
	}
}
