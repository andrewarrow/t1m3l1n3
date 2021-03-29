package main

import (
	"clt/persist"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"math/big"
)

func KeySign(msg string) (string, string) {
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, msg)
	signhash := h.Sum(nil)

	data := persist.ReadFromFile("PRIVATE_KEY")
	sDec, _ := b64.StdEncoding.DecodeString(data)
	block, _ := pem.Decode(sDec)
	x509Encoded := block.Bytes
	privatekey, _ := x509.ParseECPrivateKey(x509Encoded)

	r, s, err := ecdsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	return b64.StdEncoding.EncodeToString(r.Bytes()),
		b64.StdEncoding.EncodeToString(s.Bytes())
}
