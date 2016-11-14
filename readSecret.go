package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/RaveNoX/jenigma/getpasswd"
)

func readSecret(env string) (string, error) {
	var (
		s   string
		err error
	)

	if env != "" {
		s = os.Getenv(env)
	} else {
		r := bufio.NewReader(os.Stdin)
		s, err = getpasswd.FAsk(r, "Secret: ", os.Stderr)

		if err != nil {
			return "", fmt.Errorf("Cannot read secret from STDIN: %v", err)
		}
	}

	if s == "" {
		err = fmt.Errorf("Got enpty secret")
	}

	return s, err
}
