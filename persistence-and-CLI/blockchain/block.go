package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
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

// Serialize serializes the block
func (b *block) serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
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

// DeserializeBlock deserializes a block
func deserializeBlock(d []byte) *block {
	var block block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
