package main

import (
	"fmt"
	"os"

	"github.com/Doug-Coder64/Block_Chain/go/fs"
	"github.com/spf13/cobra"
)

const flagDataDir = "datadir"
func main() {
	var parcelCmd = &cobra.Command{
		Use: "parcel",
		Short: "Parcel App CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	parcelCmd.AddCommand(versionCmd)
	parcelCmd.AddCommand(balancesCmd())
	parcelCmd.AddCommand(runCmd())

	err := parcelCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data dir where the DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func getDataDirFromCmd(cmd *cobra.Command) string {
	dataDir, _ := cmd.Flags().GetString(flagDataDir)

	return fs.ExpandPath(dataDir)
}
