package ppk

import (
	"crypto/rsa"
	"fmt"
	"os"
	"encoding/pem"
	"crypto/x509"
	"crypto/rand"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

// This func is using for saving private and public key
func GetPrivateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error)
	}
	return privateKey
}
// Saving private key with given name..
func SavePrivateKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

// Saving public key with given name..
func SavePublicKey(fileName string, privateKey *rsa.PrivateKey) {
	pubKey := &privateKey.PublicKey
	pub, _ := ssh.NewPublicKey(pubKey)
	ioutil.WriteFile(fileName+".pub", ssh.MarshalAuthorizedKey(pub), 0777)
}


// check error
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
