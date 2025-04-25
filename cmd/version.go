package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gogogo",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("gogogo version 0.1.0")
	},
}
