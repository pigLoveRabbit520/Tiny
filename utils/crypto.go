package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"strings"
)

const (
	KEY = "R1XBXVOyZZWnAAkf"
	IV  = "h563JYDuyLfeoAam"
)

// AES-128-cbc
// AES加密数据块分组长度必须为128bit(byte[16])
// 密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个
// 加密结果最后转成base64
func En(plainText string) string {
	result, _ := encryptCBC([]byte(plainText), []byte(KEY), []byte(IV))
	return base64.StdEncoding.EncodeToString(result)
}

func De(text string) (string, error) {
	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return "", err
	}
	return decryptCBC(decodedMsg, KEY)
}

func encryptCBC(plainText []byte, key, IV []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	plainText = pad(plainText)
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	crypted := make([]byte, len(plainText))
	blockMode.CryptBlocks(crypted, plainText)
	return crypted, nil
}

func decryptCBC(decodedMsg []byte, key []byte) (plaintext string, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}

	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, []byte(IV))
	cfb.XORKeyStream(msg, msg)

	unpadMsg, err := unpad(msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}
