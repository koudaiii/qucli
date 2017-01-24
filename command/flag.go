package command

import (
	flag "github.com/spf13/pflag"
)

var (
	role       string
	visibility string
	public     bool
)

func FlagInit(args []string) error {
	flags := flag.NewFlagSet("qucli", flag.ExitOnError)

	flags.Usage = func() {
		flags.PrintDefaults()
	}

	flags.StringVar(&visibility, "visibility", "public", "visibility set to 'public' or 'private'.")
	flags.StringVar(&role, "role", "read", "role to use for the user =  ['read', 'write', 'admin'].")
	flags.BoolVar(&public, "is-public", true, "'--is-public=true' or '--is-public=false'.")

	if err := flags.Parse(args[0:]); err != nil {
		return err
	}
	return nil
}
