package main

import (
	"clt/persist"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
	"fmt"
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

}
