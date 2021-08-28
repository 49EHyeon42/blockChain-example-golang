package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"

	"github.com/49EHyeon42/blockChain-example-golang/proof-of-work/utils"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 24

// ProofOfWork represents a proof-of-work
type proofOfWork struct {
	block  *block
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
func newProofOfWork(b *block) *proofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &proofOfWork{b, target}

	return pow
}

func (pow *proofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.prevBlockHash,
			pow.block.data,
			utils.IntToHex(pow.block.timestamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *proofOfWork) run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	// fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		// fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	// fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *proofOfWork) validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
