package jsoncrypt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/RaveNoX/jenigma/jsoncrypt/pkcs7"
)

// Encrypt ecrypts unmarshaled JSON values with secret
func Encrypt(data interface{}, secret []byte) (interface{}, error) {
	var (
		err    error
		newVal interface{}
	)

	// map
	if val, ok := data.(map[string]interface{}); ok {
		var keys []string
		ret := make(map[string]interface{})

		for k := range val {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, k := range keys {
			newVal, err = Encrypt(val[k], secret)
			if err != nil {
				return nil, fmt.Errorf("Cannot encrypt %v value: %v", k, err)
			}

			ret[k] = newVal
		}

		return ret, nil
	}

	// slice
	if val, ok := data.([]interface{}); ok {
		ret := make([]interface{}, len(val))

		for i, v := range val {
			newVal, err = Encrypt(v, secret)

			if err != nil {
				return nil, fmt.Errorf("Cannot encode %v value: %v", i, err)
			}

			ret[i] = newVal
		}
		return ret, nil
	}

	// plain value
	newVal, err = encryptValue(data, secret)
	if err != nil {
		return nil, fmt.Errorf("Cannot encrypt value: %v", err)
	}
	return newVal, nil
}

func encryptValue(val interface{}, secret []byte) (string, error) {
	buff := new(bytes.Buffer)
	enc := json.NewEncoder(buff)
	err := enc.Encode(val)

	if err != nil {
		return "", fmt.Errorf("Cannot marshal value %#v to JSON, error: %v", val, err)
	}

	crypt, err := getCrypter(secret, true)

	if err != nil {
		return "", fmt.Errorf("Cannot get crypter: %v", err)
	}

	// add padding
	buffPad := pkcs7.Pad(crypt.BlockSize(), buff.Bytes())
	crypt.CryptBlocks(buffPad, buffPad)

	return base64.RawURLEncoding.EncodeToString(buffPad), nil
}
