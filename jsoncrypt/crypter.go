package jsoncrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"fmt"
)

func getCrypterKey(secret []byte) []byte {
	key := sha512.Sum512_256(secret)
	return key[:]
}

func getCrypterIV(secret []byte, blockSize int) []byte {
	buff := sha512.Sum512(secret)
	buffLen := len(buff)

	iv := make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		// position in buff
		buffIdx := i % buffLen

		// take from end
		buffIdx = buffLen - 1 - buffIdx

		iv[i] = buff[buffIdx]
	}

	return iv
}

func getCrypter(secret []byte, encrypt bool) (cipher.BlockMode, error) {
	key := getCrypterKey(secret)

	c, err := aes.NewCipher(key)

	if err != nil {
		return nil, fmt.Errorf("Cannot create AES cipher: %v", err)
	}

	iv := getCrypterIV(secret, c.BlockSize())

	if encrypt {
		return cipher.NewCBCEncrypter(c, iv), nil
	}

	return cipher.NewCBCDecrypter(c, iv), nil
}
