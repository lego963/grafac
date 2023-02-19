package main

import (
	"github.com/lego963/grafac/internal/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Terraform commands.
var Commands map[string]cli.CommandFactory

func initCommands() {
	Commands = map[string]cli.CommandFactory{
		"login": func() (cli.Command, error) {
			return &command.LoginCommand{}, nil
		},
	}
}
