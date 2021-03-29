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

/*
1. user makes keys on local hard drive
2. user proove he's "username" on twitter by posting signed message
3. user enter url of tweet
4. now we know only the person with this private key on his hard drive could have tweeted that.
5. the datetime of the tweet says who was 1st to do this


1. Upload your pub_key + username to main node

Node 1

"username" posts hi

   1. username is avail
	    assume its ok
	 2. username is in use


Node 2

*/

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

}
