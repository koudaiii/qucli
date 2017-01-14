package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/dockerepos/quay"
)

type GetCommand struct {
	Meta
}

func (c *GetCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(args[0], "/")
	if len(ss) != 3 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repos, err := quay.GetRepository(ss[1], ss[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Repository:")
	fmt.Fprintf(os.Stdout, "\tquay.io/%v/%v\n", repos.Namespace, repos.Name)

	fmt.Fprintln(os.Stdout, "Permissions:")
	permissions, err := quay.GetPermissions(ss[1], ss[2], "user")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	for _, p := range permissions.Items {
		fmt.Fprintf(os.Stdout, "\t%v(%v)\n", p.Name, p.Role)
	}

	permissions, err = quay.GetPermissions(ss[1], ss[2], "team")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	for _, p := range permissions.Items {
		fmt.Fprintf(os.Stdout, "\t%v(%v)\n", p.Name, p.Role)
	}

	return 0
}

func (c *GetCommand) Synopsis() string {
	return fmt.Sprint("Get repository and Permissions in Quay")
}

func (c *GetCommand) Help() string {
	helpText := `
dockerepos supported only Quay.io
Usage: get
  dockerepos get quay.io/koudaiii/dockerepos
`
	return strings.TrimSpace(helpText)
}
