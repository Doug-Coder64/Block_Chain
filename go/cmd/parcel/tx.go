package main

import (
	"fmt"
	"os"

	"github.com/Doug-Coder64/Block_Chain/go/database"

	"github.com/spf13/cobra"
)

const flagFrom = "from"
const flagTo = "to"
const flagValue = "value"
const flagData = "data"

func txCmd() *cobra.Command {
	var txsCmd = &cobra.Command{
		Use: "tx",
		Short: "Interact with txs (Add...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string){
		},

	}

	txsCmd.AddCommand(txAddCmd())
	return txsCmd
}

func txAddCmd() *cobra.Command {
	var txsAddCmd = &cobra.Command{
		Use: "add",
		Short: "Addsd new TX to database", 
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)
			data, _ := cmd.Flags().GetString(flagData)

			fromAcc := database.NewAccount(from)
			toAcc := database.NewAccount(to)

			tx  := database.NewTx(fromAcc, toAcc, value, data)

			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			defer state.Close()

			err = state.Add(tx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			_, err = state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println("TX successfully added to the ledger.")
		},


	}
	txsAddCmd.Flags().String(flagFrom, "", "From what account to send tokens")
	txsAddCmd.MarkFlagRequired(flagFrom)
	txsAddCmd.Flags().String(flagTo, "", "To what account to send tokens")
	txsAddCmd.MarkFlagRequired(flagTo)
	txsAddCmd.Flags().Uint(flagValue, 0, "How many tokens to send")
	txsAddCmd.MarkFlagRequired(flagValue)
	txsAddCmd.Flags().String(flagData, "", "Possible values: 'reward'")

	return txsAddCmd

}