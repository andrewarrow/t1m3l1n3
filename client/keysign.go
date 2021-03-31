package main

import (
	"clt/persist"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
	"fmt"
)

func KeySign(msg string) string {
	msgHash := sha256.New()
	msgHash.Write([]byte(msg))
	msgHashSum := msgHash.Sum(nil)

	data := persist.ReadFromFile("PRIVATE_KEY")
	block, _ := pem.Decode([]byte(data))
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return b64.StdEncoding.EncodeToString(signature)
}
