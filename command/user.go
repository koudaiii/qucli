package command

import (
	"strings"
)

type AddUserCommand struct {
	Meta
}

type DeleteUserCommand struct {
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
