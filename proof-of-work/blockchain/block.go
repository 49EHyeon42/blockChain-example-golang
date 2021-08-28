package blockchain

import (
	"fmt"
	"strconv"
	"time"
)

// Block keeps block headers
type block struct {
	timestamp     int64
	data          []byte
	prevBlockHash []byte
	hash          []byte
	nonce         int
}

func (b *block) Block() {
	fmt.Printf("Prev. hash: %x\n", b.prevBlockHash)
	fmt.Printf("Data: %s\n", b.data)
	fmt.Printf("Hash: %x\n", b.hash)
	fmt.Printf("PoW: %s\n", strconv.FormatBool(newProofOfWork(b).validate()))
	fmt.Println()
}

// NewBlock creates and returns Block
func newBlock(data string, prevBlockHash []byte) *block {
	block := &block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := newProofOfWork(block)
	nonce, hash := pow.run()

	block.hash = hash[:]
	block.nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func newGenesisBlock() *block {
	return newBlock("Genesis Block", []byte{})
}
