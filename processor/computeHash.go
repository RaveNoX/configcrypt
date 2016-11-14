package processor

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
)

func computeHash(data interface{}) (string, error) {
	buff := new(bytes.Buffer)
	err := writeJSON(data, buff)

	if err != nil {
		return "", err
	}

	summ := sha512.Sum512(buff.Bytes())
	hash := base64.RawURLEncoding.EncodeToString(summ[:])
	return hash, nil
}
