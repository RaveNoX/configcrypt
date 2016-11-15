package processor

import (
	"bufio"
	"fmt"
	"io"

	"github.com/RaveNoX/jenigma/jsoncrypt"
)

// Decrypt decrypt JSON config values
func Decrypt(secret []byte, log io.Writer, verbose bool, in io.Reader, saver Saver) error {
	var (
		err  error
		hash string
		data interface{}
	)

	reader := bufio.NewReader(in)
	if verbose {
		fmt.Fprintln(log, "Reading header")
	}

	hash, err = readHeader(reader)
	if err != nil {
		return fmt.Errorf("Cannot read header: %v", err)
	}

	if verbose {
		fmt.Fprintln(log, "Loading IN file")
	}

	data, err = readJSON(reader)
	if err != nil {
		return fmt.Errorf("Cannot parse IN file: %v", err)
	}

	if verbose {
		fmt.Fprintln(log, "Decrypting JSON values")
	}

	data, err = jsoncrypt.Decrypt(data, secret)
	if err != nil {
		return fmt.Errorf("Cannot decrypt config values, check secret")
	}

	if verbose {
		fmt.Fprintln(log, "Computing decrypted hash")
	}

	decryptedHash, err := computeHash(data)
	if err != nil {
		return fmt.Errorf("Cannot compute decrypted hash: %v", err)
	}

	if verbose {
		fmt.Fprintln(log, "Comparing hashes:")
		fmt.Fprintln(log, "  From file: ", hash)
		fmt.Fprintln(log, "  Decrypted: ", hash)
	}
	if hash != decryptedHash {
		return fmt.Errorf("Decrypted data hash not match hash from file")
	}
	fmt.Fprintln(log, "Hashes equal")

	if verbose {
		fmt.Fprintln(log, "Writing JSON data")
	}

	err = saver(func(out io.Writer) error {
		return writeJSON(data, out)
	})
	if err != nil {
		return fmt.Errorf("Cannot write JSON data: %v", err)
	}

	if verbose {
		fmt.Fprintln(log, "All done")
	}

	return nil
}
