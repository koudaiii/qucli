package command

import (
	"strings"
)

type DestroyCommand struct {
	Meta
}

func (c *DestroyCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DestroyCommand) Synopsis() string {
	return ""
}

func (c *DestroyCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
