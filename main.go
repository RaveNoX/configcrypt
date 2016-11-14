package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"bytes"

	"github.com/RaveNoX/jenigma/options"
	"github.com/RaveNoX/jenigma/processor"
)

var opts *options.Options

func init() {
	opts = &options.Options{
		Name: filepath.Base(os.Args[0]),
	}
}

func main() {
	opts.ParseOrExit(os.Args[1:])

	if opts.Verbose {
		printOpts()
	}

	if opts.Verbose {
		fmt.Fprintln(os.Stderr, "Reading secret")
	}

	secret, err := readSecret(opts.SecretEnv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read secret: %v\n", err)
		os.Exit(1)
	}

	err = process([]byte(secret))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func process(secret []byte) error {
	var (
		in  io.Reader
		out io.Writer
	)

	if opts.In == "-" {
		in = os.Stdin
	} else {
		f, err := os.Open(opts.In)
		if err != nil {
			return fmt.Errorf("Cannot open IN file: %v\n", err)
		}
		defer f.Close()
		in = f
	}

	if opts.DryRun {
		out = new(bytes.Buffer)
	} else {
		f, err := os.Create(opts.Out)
		if err != nil {
			return fmt.Errorf("Cannot create OUT file: %v\n", err)
		}
		defer f.Close()
		out = f
	}

	if opts.Encrypt {
		return processor.Encrypt(secret, os.Stderr, opts.Verbose, in, out)
	}
	return processor.Decrypt(secret, os.Stderr, opts.Verbose, in, out)
}
