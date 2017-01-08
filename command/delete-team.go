package command

import (
	"strings"
)

type DeleteTeamCommand struct {
	Meta
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
