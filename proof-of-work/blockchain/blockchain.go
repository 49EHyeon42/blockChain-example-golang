package blockchain

import "github.com/49EHyeon42/blockChain-example-golang/proof-of-work/block"

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*block.Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{block.NewGenesisBlock()}}
}
