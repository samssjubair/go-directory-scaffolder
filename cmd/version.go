package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev" // This will be set during build

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of scaffold",
	Long:  `Print the version number of scaffold`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("scaffold version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
} 