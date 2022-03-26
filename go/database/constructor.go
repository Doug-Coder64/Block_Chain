package main

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
)


func NewStateFromDisk() (*State, error) {
	//get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	
	genFilePath := filepath.Join(cwd, "database", "genesis.json")
	gen, err := loadGenesis(genFilePath)
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)

	for account, balance := range gen.Balances {
		balances[account] = balance
	}


	txDbFilePath := filepath.Join(cwd, "database", "tx.db")
	f, err := os.OpenFile(txDbFilePath, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	state := &State{balances, make([]Tx, 0), f}

	// Iterating over each tx.db file lines
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		// Convert JSON encoded TX into an object (struct)
		var tx Tx
		json.Unmarshal(scanner.Bytes(), &tx)

		/* Rebuild the state (user balances) as  series of events */
		if err := state.apply(tx); err != nil {
			return nil, err
		}
	}

	return state, nil
}