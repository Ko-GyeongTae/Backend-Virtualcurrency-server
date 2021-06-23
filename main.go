package main

import (
	"fmt"

	"github.com/Ko-GyeongTae/Backend-Nomadcoin-server/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Println(block)
	}
}
