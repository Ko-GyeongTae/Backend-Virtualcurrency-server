package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/utils"
)

const (
	signature     string = "5f8aec4b0635197bda822c695e93b6b29204af5b7e394363d0a5a84c084be98855961e840d4c5ea747a64f223135fe66c2c3904956cecfe1d5f10630fef20717"
	privateKey    string = "3077020101042077a43a9b5665699e37cbe37ee0c98d2ee418f1f3ac30cb52facfa5caad4e408ba00a06082a8648ce3d030107a14403420004c018c6400049a171aad4c2f3d15db4fbd6dd06bd5405caa892e7930524d1aeac59de02b804b47d090b880b4448153d4df152f56d70865aee9446271f0b59cb93"
	hashedMessage string = "46e89f75dec156c07a8685701e4fa023549586dab354701b270066c4d41ed5f4"
)

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)
	fmt.Printf("privateKey : %x\n", keyAsBytes)

	utils.HandleErr(err)
	message := "I love it"
	hashedMessage := utils.Hash(message)
	fmt.Println("HashedMsg : " + hashedMessage)

	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)

	signature := append(r.Bytes(), s.Bytes()...)

	fmt.Printf("Signature : %x\n", signature)
}
