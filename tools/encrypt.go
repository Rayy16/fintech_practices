package tools

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

var priKey *rsa.PrivateKey
var pubKey *rsa.PublicKey
var pubKeyString string
var priKeyString string

func InitPubKey() {
	var err error
	priKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	pubKey = &priKey.PublicKey
	pubBytes := x509.MarshalPKCS1PublicKey(pubKey)
	priBytes := x509.MarshalPKCS1PrivateKey(priKey)
	pubKeyString = string(pem.EncodeToMemory(
		&pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubBytes},
	))
	priKeyString = string(pem.EncodeToMemory(
		&pem.Block{Type: "RSA PRIVATE KEY", Bytes: priBytes},
	))
}

func Decrypt(encryptCtx string) ([]byte, error) {
	//msg, err := hex.DecodeString(encryptCtx)
	msg, err := base64.StdEncoding.DecodeString(encryptCtx)
	if err != nil {
		return nil, fmt.Errorf("hex.DecodeString err: %s", err.Error())
	}
	decryptedCtx, err := priKey.Decrypt(rand.Reader, msg, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil, fmt.Errorf("priKey.Decrypt err: %s", err.Error())
	}
	return decryptedCtx, nil
}

func Encrypt(msg string) ([]byte, error) {
	encryptData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, []byte(msg), nil)
	if err != nil {
		return nil, fmt.Errorf("rsa.EncryptOAEP err: %s", err.Error())
	}
	return encryptData, nil
}

func GetPubKey() string {
	return pubKeyString
}

func TestEncrypt() {
	InitPubKey()

	fmt.Println(pubKeyString)
	fmt.Println(priKeyString)
	encryptedTxt, err := Encrypt("cc-fintech-practices")
	if err != nil {
		fmt.Println(err.Error())
	}
	msg := base64.StdEncoding.EncodeToString(encryptedTxt)
	fmt.Println(msg)
	decryptedTxt, err := Decrypt(msg)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(decryptedTxt), string(decryptedTxt) == "cc-fintech-practices")

}
