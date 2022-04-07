package main

import (
	"fmt"
	"os"

	"github.com/Doug-Coder64/Block_Chain/go/node"
	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	var runCmd = &cobra.Command{
		Use: "run",
		Short: "Launches the parcel node and its HTTP API",
		Run: func(cmd *cobra.Command, args []string) {
			dataDir, _ := cmd.Flags().GetString(flagDataDir)

			fmt.Printf("Launching TBB node and its HTTP API...")

			err := node.Run(dataDir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	addDefaultRequiredFlags(runCmd)
	return runCmd
}