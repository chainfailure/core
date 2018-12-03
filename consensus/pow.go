package consensus

import (
	"bytes"

	"github.com/fluxchain/core/blockchain/block"
	c "github.com/fluxchain/core/crypto"
)

// Generates a proof-of-work by simply generating a blockhash and checking if
// the first 4 characters are all zeroes. Need to rework this.
func GeneratePoW(header *block.Header, target [32]byte) (c.Hash, error) {
	var result []byte
	var err error

	for {
		result, err = header.CalculateHash()
		if err != nil {
			return nil, err
		}

		if result := bytes.Compare(result, target[:]); result <= 0 {
			break
		}

		header.IncrementNonce()
	}

	return result, err
}

// Checks if the resulting block hash generated by local calculations matches
// the blockheader hash.
func ValidatePOW(header *block.Header, target [32]byte) (bool, error) {
	headerHash := header.Hash
	calculatedHash, err := header.CalculateHash()
	if err != nil {
		return false, err
	}

	return bytes.Equal(headerHash, calculatedHash), nil
}