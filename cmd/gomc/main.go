package main

import (
	"fmt"
	"gomc"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		_, err := fmt.Fprintf(os.Stderr, "No sub command specified\n")
		if err != nil {

		}
		os.Exit(1)
	}
	subCommand := os.Args[1]
	args := os.Args[2:]
	switch subCommand {
	case "dm":
		gomc.DmMain(args)
	case "httpd":
		gomc.HttpdMain(args)
	case "rest":
		gomc.RestMain(args)
	case "sqlite":
		gomc.SqliteMain(args)
	case "trial":
		gomc.TrialMain(args)
	default:
		_, err := fmt.Fprintf(os.Stderr, "No such command: %s\n", subCommand)
		if err != nil {

		}
	}
}
