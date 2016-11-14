package main

import (
	"fmt"
	"os"
)

func printOpts() {
	var mode string

	if opts.Encrypt {
		mode = "encrypt"
	} else {
		mode = "decrypt"
	}

	out := opts.Out
	if out == "-" {
		out = "<STDOUT>"
	}

	fmt.Fprintln(os.Stderr, "Mode: ", mode)
	fmt.Fprintln(os.Stderr, "In: ", opts.In)
	fmt.Fprintln(os.Stderr, "Out: ", out)
	fmt.Fprintln(os.Stderr, "Dry run: ", opts.DryRun)
	fmt.Fprintln(os.Stderr)
}
