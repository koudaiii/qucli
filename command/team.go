package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/qucli/quay"
)

type AddTeamCommand struct {
	Meta
}

type DeleteTeamCommand struct {
	Meta
}

func (c *AddTeamCommand) Run(args []string) int {
	if err := FlagInit(args); err != nil {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(args[0], "/")
	if len(ss) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repos, err := quay.AddPermission(ss[0], ss[1], "team", args[1], role, hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Added! %v(%v) in %v/%v/%v\n", repos.Name, repos.Role, hostname, ss[0], ss[1])
	return 0
}

func (c *AddTeamCommand) Synopsis() string {
	return fmt.Sprint("Add team in repository")
}

func (c *AddTeamCommand) Help() string {
	helpText := `
qucli supported only Quay.io
Usage: add-team
  qucli add-user koudaiii/qucli infrastructure --role admin
`
	return strings.TrimSpace(helpText)
}

func (c *DeleteTeamCommand) Run(args []string) int {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(args[0], "/")
	if len(ss) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	err := quay.DeletePermission(ss[0], ss[1], "team", args[1], hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Deleted! %v in %v/%v/%v\n", args[1], hostname, ss[0], ss[1])
	return 0
}

func (c *DeleteTeamCommand) Synopsis() string {
	return fmt.Sprint("Delete team in repository")
}

func (c *DeleteTeamCommand) Help() string {
	helpText := `
qucli supported only Quay.io
Usage: delete-team
  qucli delete-team koudaiii/qucli infrastructure
`
	return strings.TrimSpace(helpText)
}
