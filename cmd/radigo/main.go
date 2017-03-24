package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/Kojirou1994/radigo"
)

func main() {
	c := cli.NewCLI("radigo", radigo.Version())
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"area":        radigo.AreaCommandFactory,
		"rec":         radigo.RecCommandFactory,
		"rec-live":    radigo.RecLiveCommandFactory,
		"browse":      radigo.BrowseCommandFactory,
		"browse-live": radigo.BrowseLiveCommandFactory,
	}

	exitCode, err := c.Run()
	if err != nil {
		log.Printf("Error executing CLI: %s", err.Error())
	}

	os.Exit(exitCode)
}
