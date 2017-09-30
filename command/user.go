package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/qucli/quay"
)

type AddUserCommand struct {
	Meta
}

type DeleteUserCommand struct {
	Meta
}

func (c *AddUserCommand) Run(args []string) int {
	if err := FlagInit(args); err != nil {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	if len(subcommandArgs) < 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(subcommandArgs[0], "/")
	if len(ss) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repos, err := quay.AddPermission(ss[0], ss[1], "user", subcommandArgs[1], role, hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Added! %v(%v) in %v/%v/%v\n", repos.Name, repos.Role, hostname, ss[0], ss[1])
	return 0
}

func (c *AddUserCommand) Synopsis() string {
	return fmt.Sprint("Add user in repository")
}

func (c *AddUserCommand) Help() string {
	helpText := `
qucli supported only Quay.io
Usage: add-user
  qucli add-user koudaiii/qucli koudaiii --role admin
`
	return strings.TrimSpace(helpText)
}

func (c *DeleteUserCommand) Run(args []string) int {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(args[0], "/")
	if len(ss) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	err := quay.DeletePermission(ss[0], ss[1], "user", args[1], hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Deleted! %v in %v/%v/%v\n", args[1], hostname, ss[0], ss[1])
	return 0
}

func (c *DeleteUserCommand) Synopsis() string {
	return fmt.Sprint("Delete user in repository")
}

func (c *DeleteUserCommand) Help() string {
	helpText := `
qucli supported only Quay.io
Usage: delete-user
  qucli delete-user koudaiii/qucli koudaiii
`
	return strings.TrimSpace(helpText)
}
