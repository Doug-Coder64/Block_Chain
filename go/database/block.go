package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)


func (block *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.MyBlockHash = hash[:]
}

// New block generations
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.setHash()
	return block
}

// The Genesis block function returns first block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
