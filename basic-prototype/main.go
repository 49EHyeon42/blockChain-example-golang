package main

import (
	"github.com/49EHyeon42/blockChain-example-golang/basic-prototype/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks() {
		block.Block()
	}
}
