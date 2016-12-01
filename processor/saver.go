package processor

import (
	"io"
)

// SaveFunc function called for save
type SaveFunc func(io.Writer) error

// Saver called when save needed
type Saver func(SaveFunc) error
