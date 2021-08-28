package blockchain

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*block
}

func (bc Blockchain) Blocks() []*block {
	return bc.blocks
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := newBlock(data, prevBlock.hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block{NewGenesisBlock()}}
}
