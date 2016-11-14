// +build windows

package getpasswd

import (
	"bufio"
	"syscall"
)

// enableEnchoInput value for SetConsoleMode function:
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms686033(v=vs.85).aspx
const enableEchoInput = 0x0004

func getPassword(reader *bufio.Reader) (password string, err error) {
	var oldMode uint32

	err = syscall.GetConsoleMode(syscall.Stdin, &oldMode)
	if err != nil {
		return
	}

	newMode := (oldMode &^ enableEchoInput)

	err = setConsoleMode(syscall.Stdin, newMode)
	defer setConsoleMode(syscall.Stdin, oldMode)
	if err != nil {
		return
	}

	return readline(reader)
}

func setConsoleMode(console syscall.Handle, mode uint32) (err error) {
	dll := syscall.MustLoadDLL("kernel32")
	proc := dll.MustFindProc("SetConsoleMode")
	r, _, err := proc.Call(uintptr(console), uintptr(mode))

	if r == 0 {
		return err
	}
	return nil
}
