package tools

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

var priKey *rsa.PrivateKey
var pubKey *rsa.PublicKey
var pubKeyString string

func InitPubKey() {
	var err error
	priKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	pubKey = &priKey.PublicKey
	pubBytes := x509.MarshalPKCS1PublicKey(pubKey)
	pubBlock := pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes}
	pubKeyString = string(pem.EncodeToMemory(&pubBlock))
}

func Decrypt(encryptCtx string) ([]byte, error) {
	msg, err := hex.DecodeString(encryptCtx)
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

	encryptedTxt, err := Encrypt("cc-fintech-practices")
	if err != nil {
		fmt.Println(err.Error())
	}
	xSlice := fmt.Sprintf("%x", encryptedTxt)
	bSlice := fmt.Sprintf("%b", encryptedTxt)
	fmt.Println(GetPubKey())
	fmt.Println(encryptedTxt)
	fmt.Println(xSlice)
	fmt.Println(bSlice)

	decryptedTxt, err := Decrypt(xSlice)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(decryptedTxt), string(decryptedTxt) == "cc-fintech-practices")

}
