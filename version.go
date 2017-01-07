package main

import (
	"fmt"
)

var (
	Version  string
	Revision string
)

func printVersion() {
	fmt.Println("dockerepos version " + Version + ", build " + Revision)
}
