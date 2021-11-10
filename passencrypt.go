package main

import (
	"github.com/marspere/goencrypt"
)

func GetPlainPass(ciphPass, firstKey, secret string) (string, error) {
	cipher, err := goencrypt.NewAESCipher([]byte(firstKey), []byte(secret), goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	if err != nil {
		return "", err
	}
	return cipher.AESDecrypt(ciphPass)
}

func GetCiphPass(plainPass, firstKey, secret string) (string, error) {
	cipher, err := goencrypt.NewAESCipher([]byte(firstKey), []byte(secret), goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	if err != nil {
		return "", err
	}
	return cipher.AESEncrypt([]byte(plainPass))
}
