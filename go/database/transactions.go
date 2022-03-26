package main

import (
	"encoding/json"
	"fmt"
)

/* Add transactoins to mempool */
func (s *State) Add(tx Tx) error {
	if err := s.apply(tx); err != nil {
		return err
	}

	s.txMempool = append(s.txMempool, tx)

	return nil
}

/* Persisting transactions to disk */
func (s *State) Persist() error {

	//  making a copy of mempool 
	mempool := make([]Tx, len(s.txMempool))
	copy(mempool, s.txMempool)

	for i := 0; i < len(mempool); i++ {
		txJson, err := json.Marshal(s.txMempool[i])
		if err != nil {
			return err
		}

		if _, err = s.dbFile.Write(append(txJson, '\n')); err != nil {
			return err
		}

		s.txMempool = append(s.txMempool[:i], s.txMempool[i+1:]...)
	}

	return nil
}

/* Changing and Validating the state */
func (s *State) apply(tx Tx) error {
	if tx.isReward() {
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