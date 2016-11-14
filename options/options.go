package options

import (
	"errors"
	"fmt"
	"os"

	"github.com/juju/gnuflag"
)

var (
	// ErrorNotEnoughArguments indicates than user passed too few arguments
	ErrorNotEnoughArguments = errors.New(`Not enough arguments`)

	// ErrorTooManyArguments indicates than user passed too many arguments
	ErrorTooManyArguments = errors.New(`Too many arguments`)
)

// Options for application
type Options struct {
	Verbose, DryRun bool
	SecretEnv       string
	Encrypt         bool
	In, Out         string

	Name string
}

func (options *Options) getFlags() *gnuflag.FlagSet {
	flags := gnuflag.NewFlagSet(options.Name, gnuflag.ContinueOnError)

	flags.BoolVar(&options.DryRun, "dry", false, "do not write output, just encrypt/decrypt")
	flags.BoolVar(&options.DryRun, "d", false, "do not write output, just encrypt/decrypt")

	flags.BoolVar(&options.Verbose, "verbose", false, "be more verbose")
	flags.BoolVar(&options.Verbose, "v", false, "be more verbose")

	flags.StringVar(&options.SecretEnv, "secretenv", "", "Environment variable to read secret key from instead of STDIN")

	flags.BoolVar(&options.Encrypt, "encrypt", false, "Encrypt mode")
	flags.BoolVar(&options.Encrypt, "e", false, "Encrypt mode")

	return flags
}

// Parse parses options, emits error if any
func (options *Options) Parse(arguments []string) error {
	flags := options.getFlags()

	err := flags.Parse(false, arguments)

	if err != nil {
		return err
	}

	args := flags.Args()

	if len(args) < 2 {
		return ErrorNotEnoughArguments
	}

	if len(args) > 2 {
		return ErrorTooManyArguments
	}

	options.In = args[0]
	options.Out = args[1]

	return nil
}

func (options *Options) printUsage() {
	flags := options.getFlags()

	format := "%s\n    %s\n"
	fmt.Fprintf(os.Stderr, "Usage: %s: [args] <in> <out>\n", options.Name)
	fmt.Fprintf(os.Stderr, format, "<in>", `Path to original file`)
	fmt.Fprintf(os.Stderr, format, "<out>", `Path to output file, "-" for STDOUT`)
	flags.PrintDefaults()
}

// ParseOrExit parses options, if any error calls os.Exit(2)
func (options *Options) ParseOrExit(arguments []string) {
	err := options.Parse(arguments)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n", err)
		options.printUsage()
		os.Exit(2)
	}
}
