package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/qcli/quay"
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
	if len(ss) != 3 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repos, err := quay.AddPermission(ss[1], ss[2], "team", args[1], role)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Added! %v(%v) in quay.io/%v/%v\n", repos.Name, repos.Role, ss[1], ss[2])
	return 0
}

func (c *AddTeamCommand) Synopsis() string {
	return fmt.Sprint("Add team in repository")
}

func (c *AddTeamCommand) Help() string {
	helpText := `
qcli supported only Quay.io
Usage: add-team
  qcli add-user quay.io/koudaiii/qcli infrastructure --role admin
`
	return strings.TrimSpace(helpText)
}

func (c *DeleteTeamCommand) Run(args []string) int {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(args[0], "/")
	if len(ss) != 3 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	err := quay.DeletePermission(ss[1], ss[2], "team", args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Deleted! %v in quay.io/%v/%v\n", args[1], ss[1], ss[2])
	return 0
}

func (c *DeleteTeamCommand) Synopsis() string {
	return fmt.Sprint("Delete team in repository")
}

func (c *DeleteTeamCommand) Help() string {
	helpText := `
qcli supported only Quay.io
Usage: delete-team
  qcli delete-team quay.io/koudaiii/qcli infrastructure
`
	return strings.TrimSpace(helpText)
}
