package database

import "os"

//Block contains info
type Block struct {
	Timestamp int64
	PreviousBlockHash []byte
	MyBlockHash []byte
	AllData []byte
}

type Blockchain struct {
	Blocks []*Block
}

type Account string

type Tx struct {
	From 	Account	`json: "from"`
	To 		Account 	`json: "to"`
	Value 	uint	`json: "value"`
	Data 	string 	`json: "data"`
}

type State struct {
	Balances	map[Account]uint
	txMempool   []Tx

	dbFile *os.File
}