package command

import (
	"strings"
)

type AddUserCommand struct {
	Meta
}

func (c *AddUserCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *AddUserCommand) Synopsis() string {
	return ""
}

func (c *AddUserCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
