package blockchain

import (
	"bytes"
	"crypto/sha256"
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
}

func (b block) Block() {
	fmt.Printf("Prev. hash: %x\n", b.prevBlockHash)
	fmt.Printf("Data: %s\n", b.data)
	fmt.Printf("Hash: %x\n", b.hash)
	fmt.Println()
}

// SetHash calculates and sets block hash
func (b *block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	headers := bytes.Join([][]byte{b.prevBlockHash, b.data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.hash = hash[:]
}

// NewBlock creates and returns Block
func newBlock(data string, prevBlockHash []byte) *block {
	block := &block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.setHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *block {
	return newBlock("Genesis Block", []byte{})
}
