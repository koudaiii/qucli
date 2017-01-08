package command

import (
	"strings"
)

type DeleteUserCommand struct {
	Meta
}

func (c *DeleteUserCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DeleteUserCommand) Synopsis() string {
	return ""
}

func (c *DeleteUserCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
