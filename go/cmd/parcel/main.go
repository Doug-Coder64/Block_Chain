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

	err := parcelCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}