package command

import (
	flag "github.com/spf13/pflag"
)

var (
	subcommandArgs []string
	role           string
	visibility     string
	public         bool
	hostname       string
)

func FlagInit(args []string) error {
	flags := flag.NewFlagSet("qucli", flag.ExitOnError)

	flags.Usage = func() {
		flags.PrintDefaults()
	}

	flags.StringVar(&visibility, "visibility", "public", "visibility set to 'public' or 'private'.")
	flags.StringVar(&role, "role", "read", "role to use for the user =  ['read', 'write', 'admin'].")
	flags.BoolVar(&public, "is-public", true, "'--is-public=true' or '--is-public=false'.")
	flags.StringVar(&hostname, "hostname", "quay.io", "If you use enterprise plan, set hostname option. ex '--hostname=quay.example.com'")

	if err := flags.Parse(args[0:]); err != nil {
		return err
	}
	subcommandArgs = flags.Args()
	return nil
}
