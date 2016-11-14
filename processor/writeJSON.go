package processor

import (
	"encoding/json"
	"fmt"
	"io"
)

func writeJSON(data interface{}, writer io.Writer) error {
	enc := json.NewEncoder(writer)
	enc.SetIndent("", "  ")

	err := enc.Encode(data)

	if err != nil {
		return fmt.Errorf("Cannot write JSON: %v", err)
	}

	return nil
}
