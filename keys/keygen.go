package keys

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"t1m3l1n3/persist"
)

func KeyGen() (string, string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})

	pubkey := &privateKey.PublicKey

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(pubkey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}
func KeyGenSave(priv, pub string) {
	location := persist.UserHomeDir()
	persist.SaveToFile("PRIVATE_KEY", priv)
	fmt.Println("Private Key Saved To: ", fmt.Sprintf("%s/%s/%s", location, persist.DIRNAME, "PRIVATE_KEY"))

	persist.SaveToFile("PUBLIC_KEY", pub)
	fmt.Println("Public Key Saved To: ", fmt.Sprintf("%s/%s/%s", location, persist.DIRNAME, "PUBLIC_KEY"))

	fmt.Printf("\n\n")
	fmt.Printf("-------------------\n")
	fmt.Printf("COPY AND PASTE DATA BELOW AND SEND YOURSELF AN EMAIL \n")
	fmt.Printf("THERE IS NO FORGET PASSWORD LINK THIS IS IT \n")
	fmt.Printf("-------------------\n")
	fmt.Printf("\n\n")
	fmt.Printf("%s\n", priv)
	fmt.Printf("%s\n", pub)
	fmt.Printf("\n\n")

	//DoTestSignAndVerify()
}

func DoTestSignAndVerify() {
	msg := "hi"
	pub := persist.ReadFromFile("PUBLIC_KEY")
	fmt.Println(pub)
	blockPub, _ := pem.Decode([]byte(pub))
	genericPublicKey, _ := x509.ParsePKIXPublicKey(blockPub.Bytes)
	publicKey := genericPublicKey.(*rsa.PublicKey)

	data := persist.ReadFromFile("PRIVATE_KEY")
	s := KeySign(data, msg)

	msgHash := sha256.New()
	msgHash.Write([]byte(msg))
	msgHashSum := msgHash.Sum(nil)

	valid := rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, []byte(s), nil)
	fmt.Println(valid)

}
