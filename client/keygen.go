package main

import (
	"clt/persist"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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

	key := fmt.Sprintf("%x\n%x\n%x\n", privatekey.X, privatekey.Y, privatekey.D)
	location := persist.UserHomeDir()
	persist.SaveToFile("PRIVATE_KEY", key)
	fmt.Println("Private Key Saved To: ", fmt.Sprintf("%s/%s/%s", location, persist.DIRNAME, "PRIVATE_KEY"))

}
