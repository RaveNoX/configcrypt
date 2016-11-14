package processor

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/RaveNoX/go-jsoncommentstrip"
	"github.com/spkg/bom"
)

func readJSON(reader io.Reader) (interface{}, error) {
	var data interface{}

	bomReader := bom.NewReader(reader)
	jsonCommentReader := jsoncommentstrip.NewReader(bomReader)

	dec := json.NewDecoder(jsonCommentReader)
	dec.UseNumber()

	err := dec.Decode(&data)

	if err != nil {
		return nil, fmt.Errorf("Cannot parse JSON: %v", err)
	}

	return data, nil
}
