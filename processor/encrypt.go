package processor

import (
	"fmt"
	"io"

	"github.com/RaveNoX/jenigma/jsoncrypt"
)

// Encrypt encrypt JSON config values
func Encrypt(secret []byte, log io.Writer, verbose bool, in io.Reader, out io.Writer) error {
	var (
		err  error
		hash string
		data interface{}
	)

	if verbose {
		fmt.Fprintln(log, "Loading IN file")
	}

	data, err = readJSON(in)
	if err != nil {
		return fmt.Errorf("Cannot parse IN file: %v", err)
	}

	if verbose {
		fmt.Fprintln(log, "Computing hash")
	}

	hash, err = computeHash(data)
	if err != nil {
		return fmt.Errorf("Cannot compute IN hash: %v", err)
	}

	if verbose {
		fmt.Fprintf(log, "IN hash: %s\n", hash)
		fmt.Fprintln(log, "Encrypting values")
	}

	data, err = jsoncrypt.Encrypt(data, secret)
	if err != nil {
		return fmt.Errorf("Cannot encrypt config values: %v", err)
	}

	if verbose {
		fmt.Fprintln(log, "Writing header")
	}

	err = writeHeader(hash, out)
	if err != nil {
		return fmt.Errorf("Cannot write header: %v", err)
	}

	if verbose {
		fmt.Fprintln(log, "Writing JSON data")
	}

	err = writeJSON(data, out)
	if err != nil {
		return fmt.Errorf("Cannot write JSON data: %v", err)
	}

	if verbose {
		fmt.Fprintln(log, "All done")
	}

	return nil
}
