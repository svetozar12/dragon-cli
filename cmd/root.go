package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/svetozar12/dragon-cli/cmd/generate"
)

var rootCmd = &cobra.Command{
	Use:   "dragon-cli",
	Short: "Generate full-stack projects based on user choices",
	Run: func(cmd *cobra.Command, args []string) {
		generate.Generate()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
