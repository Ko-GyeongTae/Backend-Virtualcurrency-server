package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/utils"
)

const (
	signature     string = "5f8aec4b0635197bda822c695e93b6b29204af5b7e394363d0a5a84c084be98855961e840d4c5ea747a64f223135fe66c2c3904956cecfe1d5f10630fef20717"
	privateKey    string = "3077020101042077a43a9b5665699e37cbe37ee0c98d2ee418f1f3ac30cb52facfa5caad4e408ba00a06082a8648ce3d030107a14403420004c018c6400049a171aad4c2f3d15db4fbd6dd06bd5405caa892e7930524d1aeac59de02b804b47d090b880b4448153d4df152f56d70865aee9446271f0b59cb93"
	hashedMessage string = "46e89f75dec156c07a8685701e4fa023549586dab354701b270066c4d41ed5f4"
)

func Start() {
	privByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	privateKey, err := x509.ParseECPrivateKey(privByte)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
	fmt.Println(bigR, bigS)
}
