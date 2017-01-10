package command

import (
"strings"
"fmt"
)

type GetCommand struct {
	Meta
}

func (c *GetCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GetCommand) Synopsis() string {
	return fmt.Sprint("Get repository and Permissions in Quay")
}

func (c *GetCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
