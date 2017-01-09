package command

import (
	"strings"
)

type AddTeamCommand struct {
	Meta
}

type DeleteTeamCommand struct {
	Meta
}


func (c *AddTeamCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *AddTeamCommand) Synopsis() string {
	return ""
}

func (c *AddTeamCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

func (c *DeleteTeamCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DeleteTeamCommand) Synopsis() string {
	return ""
}

func (c *DeleteTeamCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
