package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mksipp",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mksipp v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
