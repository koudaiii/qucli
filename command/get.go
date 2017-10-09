package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/koudaiii/qucli/quay"
)

type GetCommand struct {
	Meta
}

func (c *GetCommand) Run(args []string) int {
	if err := FlagInit(args); err != nil {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	if len(subcommandArgs) != 1 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	ss := strings.Split(subcommandArgs[0], "/")
	if len(ss) != 2 {
		fmt.Fprintln(os.Stderr, c.Help())
		os.Exit(1)
	}

	repos, err := quay.GetRepository(ss[0], ss[1], hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "Repository:")
	fmt.Fprintf(os.Stdout, "\t%v/%v/%v\n", hostname, repos.Namespace, repos.Name)

	fmt.Fprintln(os.Stdout, "Visibility:")
	if repos.IsPublic == true {
		fmt.Fprintln(os.Stdout, "\tpublic")
	} else {
		fmt.Fprintln(os.Stdout, "\tprivate")
	}
	fmt.Fprintln(os.Stdout, "Permissions:")

	permissions, err := quay.GetPermissions(ss[0], ss[1], "user", hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	for _, p := range permissions.Items {
		fmt.Fprintf(os.Stdout, "\t%v(%v)\n", p.Name, p.Role)
	}

	permissions, err = quay.GetPermissions(ss[0], ss[1], "team", hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	for _, p := range permissions.Items {
		fmt.Fprintf(os.Stdout, "\t%v(%v)\n", p.Name, p.Role)
	}

	fmt.Fprintln(os.Stdout, "Notifications:")

	notitications, err := quay.ListRepositoryNotifications(ss[0], ss[1], hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "\tTitle\tEvent\tMethod\tEventConfig\tUUID\tNumberOfFailures\tConfig")

	for _, n := range notitications.Items {
		fmt.Fprintf(os.Stdout, "\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", n.Title, n.Event, n.Method, n.EventConfig, n.UUID, n.NumberOfFailures, n.Config)
	}
	return 0
}

func (c *GetCommand) Synopsis() string {
	return fmt.Sprint("Get repository and Permissions in Quay")
}

func (c *GetCommand) Help() string {
	helpText := `
qucli supported only Quay.io
Usage: get
  qucli get koudaiii/qucli
`
	return strings.TrimSpace(helpText)
}
