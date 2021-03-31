package main

import (
	"clt/persist"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func KeyGen() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})

	location := persist.UserHomeDir()
	persist.SaveToFile("PRIVATE_KEY", string(pemEncoded))
	fmt.Println("Private Key Saved To: ", fmt.Sprintf("%s/%s/%s", location, persist.DIRNAME, "PRIVATE_KEY"))

	pubkey := &privateKey.PublicKey

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(pubkey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509EncodedPub})
	persist.SaveToFile("PUBLIC_KEY", string(pemEncodedPub))
	fmt.Println("Public Key Saved To: ", fmt.Sprintf("%s/%s/%s", location, persist.DIRNAME, "PUBLIC_KEY"))

	fmt.Printf("\n\n")
	fmt.Printf("-------------------\n")
	fmt.Printf("COPY AND PASTE DATA BELOW AND SEND YOURSELF AN EMAIL \n")
	fmt.Printf("THERE IS NO FORGET PASSWORD LINK THIS IS IT \n")
	fmt.Printf("-------------------\n")
	fmt.Printf("\n\n")
	fmt.Printf("%s\n", string(pemEncoded))
	fmt.Printf("%s\n", string(pemEncodedPub))
	fmt.Printf("\n\n")

	DoTestSignAndVerify()
}

func DoTestSignAndVerify() {
	msg := "hi"
	pub := persist.ReadFromFile("PUBLIC_KEY")
	fmt.Println(pub)
	blockPub, _ := pem.Decode([]byte(pub))
	genericPublicKey, _ := x509.ParsePKIXPublicKey(blockPub.Bytes)
	publicKey := genericPublicKey.(*rsa.PublicKey)

	s := KeySign(msg)
	fmt.Println(s)

	msgHash := sha256.New()
	msgHash.Write([]byte(msg))
	msgHashSum := msgHash.Sum(nil)

	valid := rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, []byte(s), nil)
	fmt.Println(valid)

}
