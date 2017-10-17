package command

import (
	flag "github.com/spf13/pflag"
)

var (
	subcommandArgs []string
	role           string
	visibility     string
	public         bool
	event          string
	method         string
	url            string
	email          string
	title          string
	hostname       string
	ref            string
	level          string
)

func FlagInit(args []string) error {
	flags := flag.NewFlagSet("qucli", flag.ExitOnError)

	flags.Usage = func() {
		flags.PrintDefaults()
	}

	flags.StringVar(&visibility, "visibility", "public", "visibility set to 'public' or 'private'.")
	flags.StringVar(&role, "role", "read", "role to use for the user =  ['read', 'write', 'admin'].")
	flags.BoolVar(&public, "is-public", true, "'--is-public=true' or '--is-public=false'.")
	flags.StringVar(&event, "event", "", "set 'evnet'.  ['repo_push', 'build_queued', 'build_start', 'build_success', 'build_failure', 'build_cancelled', 'vulnerability_found'].")
	flags.StringVar(&level, "level", "", "if you use 'vulnerability_found' method, A vulnerability must have a severity of the chosen level (highest level is 0).[0-6]")
	flags.StringVar(&ref, "ref", "", "if you use event excluding 'repo_push' event, an optional regular expression for matching the git branch or tag git ref. If left blank, the notification will fire for all builds.(refs/heads/somebranch)|(refs/tags/sometag)")
	flags.StringVar(&method, "method", "", "set 'method'.  ['webhook', 'slack', 'email'].")
	flags.StringVar(&email, "email", "", "if you use 'email' method, set E-mail address. 'test@example.com'.")
	flags.StringVar(&url, "url", "", "if you use 'webhook' or 'slack' method, set url. 'http://url/goes/here' or 'https://hooks.slack.com/service/{some}/{token}/{here}'.")
	flags.StringVar(&title, "title", "", "The title for a notification is an optional field for a human-readable title for the notification.")
	flags.StringVar(&hostname, "hostname", "quay.io", "If you use enterprise plan, set hostname option. ex '--hostname=quay.example.com'")

	if err := flags.Parse(args[0:]); err != nil {
		return err
	}
	subcommandArgs = flags.Args()
	return nil
}
