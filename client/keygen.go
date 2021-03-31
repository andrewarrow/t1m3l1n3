package main

import (
	"clt/persist"
	"crypto/ecdsa"
	"crypto/elliptic"
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

func KeyGen() {
	pubkeyCurve := elliptic.P256()

	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	x509Encoded, _ := x509.MarshalECPrivateKey(privatekey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	key := b64.StdEncoding.EncodeToString(pemEncoded)

	location := persist.UserHomeDir()
	persist.SaveToFile("PRIVATE_KEY", key)
	fmt.Println("Private Key Saved To: ", fmt.Sprintf("%s/%s/%s", location, persist.DIRNAME, "PRIVATE_KEY"))

	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(pubkey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	keyPub := b64.StdEncoding.EncodeToString(pemEncodedPub)
	persist.SaveToFile("PUBLIC_KEY", keyPub)
	fmt.Println("Public Key Saved To: ", fmt.Sprintf("%s/%s/%s", location, persist.DIRNAME, "PUBLIC_KEY"))

	fmt.Printf("\n\n")
	fmt.Printf("-------------------\n")
	fmt.Printf("COPY AND PASTE DATA BELOW AND SEND YOURSELF AN EMAIL \n")
	fmt.Printf("THERE IS NO FORGET PASSWORD LINK THIS IS IT \n")
	fmt.Printf("-------------------\n")
	fmt.Printf("\n\n")
	fmt.Printf("%s\n", key)
	fmt.Printf("%s\n", keyPub)
	fmt.Printf("\n\n")

	DoTestSignAndVerify()
}

func DoTestSignAndVerify() {
	msg := "hi"
	pub := persist.ReadFromFile("PUBLIC_KEY")
	fmt.Println(pub)
	sDec, _ := b64.StdEncoding.DecodeString(pub)
	fmt.Println(sDec)

	fmt.Println(pub)
	blockPub, e := pem.Decode(sDec)
	fmt.Println(e)
	fmt.Println(blockPub)
	x509EncodedPub := blockPub.Bytes
	fmt.Println(blockPub.Bytes)
	genericPublicKey, ee := x509.ParsePKIXPublicKey(x509EncodedPub)
	fmt.Println(genericPublicKey)
	fmt.Println(ee)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	r, s := KeySign(msg)
	fmt.Println(r)
	fmt.Println(s)

	var h hash.Hash
	h = md5.New()
	io.WriteString(h, msg)
	signhash := h.Sum(nil)
	bigr := big.NewInt(0)
	rDec, _ := b64.StdEncoding.DecodeString(r)
	bigr = bigr.SetBytes(rDec)
	bigs := big.NewInt(0)
	sDec2, _ := b64.StdEncoding.DecodeString(s)
	bigs = bigs.SetBytes(sDec2)
	verifystatus := ecdsa.Verify(publicKey, signhash, bigr, bigs)
	fmt.Println(verifystatus)
}
