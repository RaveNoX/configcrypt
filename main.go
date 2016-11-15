package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

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

func Saver(saveFn processor.SaveFunc) error {
	if opts.DryRun {
		return saveFn(new(bytes.Buffer))
	}

	if opts.Out == "-" {
		return saveFn(os.Stdout)
	}

	f, err := os.Create(opts.Out)
	if err != nil {
		return fmt.Errorf("Cannot create OUT file: %v\n", err)
	}
	defer f.Close()

	return saveFn(f)
}

func process(secret []byte) error {
	in, err := os.Open(opts.In)
	if err != nil {
		return fmt.Errorf("Cannot open IN file: %v\n", err)
	}
	defer in.Close()

	if opts.Encrypt {
		return processor.Encrypt(secret, os.Stderr, opts.Verbose, in, Saver)
	}
	return processor.Decrypt(secret, os.Stderr, opts.Verbose, in, Saver)
}
