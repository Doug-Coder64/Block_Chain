package main

import (
	"fmt"

	"github.com/spf13/cobra"
)


const Major = "0"
const Minor = "1"
const Fix = "0"
const Verbal = "TX ADD && Balances list"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Describes Version.", 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Version: %s.%s.%s-beta %s", Major, Minor, Fix, Verbal))
	},
}
