package getpasswd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Ask the user to enter a password with input hidden.

Prompt is a string to display before the user's input.

Returns the provided password, or an error if the command failed.
*/
func Ask(reader *bufio.Reader, prompt string) (password string, err error) {
	return FAsk(reader, prompt, os.Stdout)
}

/*
FAsk is the same as Ask, except it is possible to specify the file to write
the prompt to. If 'nil' is passed as the writer, no prompt will be written.
*/
func FAsk(reader *bufio.Reader, prompt string, wr io.Writer) (password string, err error) {
	if wr != nil && prompt != "" {
		fmt.Fprint(wr, prompt) // Display the prompt.
	}
	password, err = getPassword(reader)

	// Carriage return after the user input.
	if wr != nil {
		fmt.Fprintln(wr, "")
	}
	return
}

func readline(reader *bufio.Reader) (value string, err error) {
	var (
		r    rune
		buff []rune
	)

	for {
		r, _, err = reader.ReadRune()

		if err != nil && err != io.EOF {
			break
		}

		if r == '\n' {
			break
		}

		buff = append(buff, r)
	}

	str := string(buff)
	str = strings.TrimSuffix(str, "\r")

	return str, err
}
