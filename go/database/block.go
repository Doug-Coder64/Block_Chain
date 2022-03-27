package database

import (
	"crypto/sha256"
	"encoding/json"
)

type Hash [32]byte

type Block struct {
	Header BlockHeader
	TXs []Tx
}

type BlockHeader struct {
	Parent Hash
	Time uint64
}

func (b Block) Hash() (Hash, error) {
	blockJson, err := json.Marshal(b)
	if err != nil {
		return Hash{}, err
	}
	return sha256.Sum256(blockJson), nil
}