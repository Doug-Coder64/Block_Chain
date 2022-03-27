package database

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Snapshot [32]byte

type State struct {
	Balances	map[Account]uint
	txMempool   []Tx

	dbFile *os.File
	snapshot Snapshot
}


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
	state := &State{balances, make([]Tx, 0), f, Snapshot{}}

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

func (s *State) LatestSnapshot() Snapshot {
	return s.snapshot
}

func (s *State) Close() error {
	return s.dbFile.Close()
}

/* Add transactoins to mempool */
func (s *State) Add(tx Tx) error {
	if err := s.apply(tx); err != nil {
		return err
	}

	s.txMempool = append(s.txMempool, tx)

	return nil
}

/* Persisting transactions to disk */
func (s *State) Persist() (Snapshot, error) {

	//  making a copy of mempool 
	mempool := make([]Tx, len(s.txMempool))
	copy(mempool, s.txMempool)

	for i := 0; i < len(mempool); i++ {
		txJson, err := json.Marshal(s.txMempool[i])
		if err != nil {
			return Snapshot{}, err
		}

		if _, err = s.dbFile.Write(append(txJson, '\n')); err != nil {
			return Snapshot{}, err
		}

		s.txMempool = append(s.txMempool[:i], s.txMempool[i+1:]...)
	}

	return s.snapshot, nil
}

/* Changing and Validating the state */
func (s *State) apply(tx Tx) error {
	if tx.IsReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}

	if tx.Value > s.Balances[tx.From]{
		return fmt.Errorf("insufficient_balance")
	}

	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value

	return nil
}

/* New 'snapshot' for every transaction */
func (s *State) doSnapshot() error {
	// Re-read whole file from the first byte
	_, err := s.dbFile.Seek(0,0)
	if err != nil {
		return err
	}

	txsData, err := ioutil.ReadAll(s.dbFile)
	if err != nil {
		return err
	}

	s.snapshot = sha256.Sum256(txsData)
	return nil
}