package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

type myCipherClient interface {
	GetAESDecrypted(encrypted string) ([]byte, error)
	GetAESEncrypted(plaintext string) (string, error)
}

type MyChiper struct {
	Key []byte
	Iv  []byte
}

func NewCipherClient(key, iv string) myCipherClient {
	return &MyChiper{
		Key: []byte(key),
		Iv:  []byte(iv),
	}
}

func (c MyChiper) GetAESDecrypted(encrypted string) ([]byte, error) {

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(c.Key)

	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, c.Iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = c.PKCS5UnPadding(ciphertext)

	return ciphertext, nil
}

func (c MyChiper) PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

func (c MyChiper) GetAESEncrypted(plaintext string) (string, error) {

	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher(c.Key)

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, c.Iv)
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil
}
