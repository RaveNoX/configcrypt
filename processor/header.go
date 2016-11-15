package processor

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const headerPrefix = `//jenigma:`

func writeHeader(hash string, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s%s\n", headerPrefix, hash)
	return err
}

func readHeader(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')

	if err != nil {
		return "", err
	}

	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, headerPrefix) {
		return line, fmt.Errorf("Cannot find header")
	}

	// extract hash
	line = line[len(headerPrefix):]

	if len(line) == 0 {
		return line, fmt.Errorf("Invalid hash size")
	}

	return line, nil
}
