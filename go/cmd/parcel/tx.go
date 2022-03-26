package main

import (
	"github.com/spf13/cobra"
)

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
			value, _ := cmd.Flags().Getuint(flagValue)

			fromAcc := database.NewAccount(from)
			toAcc := database.NewAccount(to)
		}

	}
}