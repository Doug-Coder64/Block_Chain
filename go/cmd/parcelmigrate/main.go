package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Doug-Coder64/Block_Chain/go/database"
)

func main() {
	state, err := database.NewStateFromDisk()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	block0 := database.NewBlock(
		database.Hash{},
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("Doug", "Doug", 3, ""),
			database.NewTx("Doug", "Doug", 700, "reward"),
		},
	)

	state.AddBlock(block0)
	block0hash, _ := state.Persist()

	block1 := database.NewBlock(
		block0hash,
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("Doug", "babayaga", 2000, ""),
			database.NewTx("doug", "doug", 100, "reward"),
			database.NewTx("babayaga", "doug", 1, ""),
			database.NewTx("babayaga", "caesar", 1000, ""),
			database.NewTx("babayaga", "doug", 50, ""),
			database.NewTx("doug", "doug", 600, "reward"),
		},
	)

	state.AddBlock(block1)
	state.Persist()
}