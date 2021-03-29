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
	"os"
)

func KeySign() {
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, "This is a message to be signed!")
	signhash := h.Sum(nil)

	data := persist.ReadFromFile("PRIVATE_KEY")
	sDec, _ := b64.StdEncoding.DecodeString(data)
	block, _ := pem.Decode(sDec)
	x509Encoded := block.Bytes
	privatekey, _ := x509.ParseECPrivateKey(x509Encoded)

	r, s, err := ecdsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	sEnc := b64.StdEncoding.EncodeToString(signature)

	fmt.Printf("Signature : %s\n", sEnc)

	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey
	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus)
}
