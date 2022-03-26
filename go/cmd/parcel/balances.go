package main

import (
	"fmt"
	"os"

	"github.com/Doug-Coder64/Block_Chain/go/database"
	"github.com/spf13/cobra"
)
func balanceCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use: "balances",
		Short: "interact with balances (list...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil // incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}

	balancesCmd.AddCommand(balancesListCmd())

	return balancesCmd
}

func balancesListCmd() *cobra.Command{
	var balancesListCmd = &cobra.Command{
	Use: "list", 
	Short: "Lists all balances.",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := database.NewStateFromDisk()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		defer state.Close()

		fmt.Println("Acocunts Balances:")
		fmt.Println("__________________")
		fmt.Println("")
		for account, balances := range state.Balances {
			fmt.Println(fmt.Sprintf("%s: %d", account, balances))
		}
	},
	}

	return balancesListCmd
}