package main

import (
	"fmt"
	"os"

	"github.com/Doug-Coder64/Block_Chain/go/database"
	"github.com/spf13/cobra"
)

func balancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances (list...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	balancesCmd.AddCommand(balancesListCmd)

	return balancesCmd
}

var balancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all balances.",
	Run: func(cmd *cobra.Command, args []string) {
		dataDir, _ := cmd.Flags().GetString(flagDataDir)
		state, err := database.NewStateFromDisk(dataDir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer state.Close()

		fmt.Printf("Accounts balances at %x:\n", state.LatestBlockHash())
		fmt.Println("__________________")
		fmt.Println("")
		for account, balance := range state.Balances {
			fmt.Println(fmt.Sprintf("%s: %d", account, balance))
		}
	},

}