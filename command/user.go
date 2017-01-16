package command

import (
	"fmt"
	"strings"
	"os"

	"github.com/koudaiii/dockerepos/quay"
	flag "github.com/spf13/pflag"
)

var role string

type AddUserCommand struct {
	Meta
}

type DeleteUserCommand struct {
	Meta
}

func (c *AddUserCommand) Run(args []string) int {
	flags := flag.NewFlagSet("dockerepos", flag.ExitOnError)

	flags.Usage = func() {
		flags.PrintDefaults()
	}

	flags.StringVar(&role, "role", "read", "role to use for the user =  ['read', 'write', 'admin']. default: read")

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(args[0], "/")
	if len(ss) != 3 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repos, err := quay.AddPermission(ss[1], ss[2],"user",args[1],role)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Added! %v(%v) in quay.io/%v/%v\n", repos.Name, repos.Role,ss[1], ss[2])
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
