package lib

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

const (
	DESKEY = "SHANGZUO"
	DESIV  = "HERE20IV"
)

func Encrypt(plainText []byte) (string, error) {
	key := []byte(DESKEY)
	iv := []byte(DESIV)
	block, err := des.NewCipher(key)

	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.URLEncoding.EncodeToString(cryted), nil
}

func Decrypt(encodeText string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(encodeText)
	if err != nil {
		return "", err
	}
	key := []byte(DESKEY)
	iv := []byte(DESIV)

	block, err := des.NewCipher(key)

	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
