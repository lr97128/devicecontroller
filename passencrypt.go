package main

import (
	"github.com/marspere/goencrypt"
)

func GetPlainPass(ciphPass, secret string) (string, error) {
	var firstKey = "0123456789asdfgh"
	cipher, err := goencrypt.NewAESCipher([]byte(firstKey), []byte(secret), goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	if err != nil {
		return "", err
	}
	plainPass, err := cipher.AESDecrypt(ciphPass)
	if err != nil {
		return "", err
	}
	return string(plainPass), nil
}
