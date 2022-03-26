package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
func main() {
	var parcelCmd = &cobra.Command{
		Use: "parcel",
		Short: "Parcel App CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	parcelCmd.AddCommand(versionCmd)
	parcelCmd.AddCommand(balancesCmd())
	parcelCmd.AddCommand(txCmd())

	err := parcelCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}