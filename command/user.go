package command

import (
	"strings"
	"fmt"
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
	return fmt.Sprint("Add user in repository")
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
	return fmt.Sprint("Delete user in repository")
}

func (c *DeleteUserCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
