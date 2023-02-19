package command

import (
	"strings"
)

type LoginCommand struct {
}

func (c *LoginCommand) Run(rawArgs []string) int {}

func (c *LoginCommand) Help() string {
	helpText := `
bla-bla
`
	return strings.TrimSpace(helpText)
}

func (c *LoginCommand) Synopsis() string {
	return "Bla-bla"
}
