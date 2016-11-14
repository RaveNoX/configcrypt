package jsoncrypt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/RaveNoX/jenigma/jsoncrypt/pkcs7"
)

// Decrypt decrypts unmarshaled JSON values with secret
func Decrypt(data interface{}, secret []byte) (interface{}, error) {
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
			newVal, err = Decrypt(val[k], secret)
			if err != nil {
				return nil, fmt.Errorf("Cannot decode %v value: %v", k, err)
			}

			ret[k] = newVal
		}

		return ret, nil
	}

	// slice
	if val, ok := data.([]interface{}); ok {
		ret := make([]interface{}, len(val))

		for i, v := range val {
			newVal, err = Decrypt(v, secret)

			if err != nil {
				return nil, fmt.Errorf("Cannot decode %v value: %v", i, err)
			}

			ret[i] = newVal
		}
		return ret, nil
	}

	// plain value
	newVal, err = decryptValue(data, secret)
	if err != nil {
		return nil, fmt.Errorf("Cannot decode value: %v", err)
	}
	return newVal, nil
}

func decryptValue(val interface{}, secret []byte) (interface{}, error) {
	strVal, ok := val.(string)

	if !ok {
		// keep as-is
		return val, nil
	}

	crypt, err := getCrypter(secret, false)

	if err != nil {
		return "", fmt.Errorf("Cannot get crypter: %v", err)
	}

	buff, err := base64.RawURLEncoding.DecodeString(strVal)

	if err != nil {
		return nil, fmt.Errorf("Cannot decrypt value %v: %v", val, err)
	}

	buffLen := len(buff)
	if buffLen%crypt.BlockSize() != 0 {
		return nil, fmt.Errorf("Len of value does invalid for crypter: %v %% %v = %v", buffLen, crypt.BlockSize(), buffLen%crypt.BlockSize())
	}

	crypt.CryptBlocks(buff, buff)

	// remove padding
	buffPad, err := pkcs7.Unpad(crypt.BlockSize(), buff)
	if err != nil {
		// looks like no padding
		buffPad = buff
	}

	dec := json.NewDecoder(bytes.NewReader(buffPad))
	dec.UseNumber()

	var data interface{}

	err = dec.Decode(&data)

	if err != nil {
		return nil, fmt.Errorf("Cannot decode JSON from value \"%s\",  error: %v", string(buff), err)
	}

	return data, nil
}
