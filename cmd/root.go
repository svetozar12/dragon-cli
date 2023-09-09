package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/svetozar12/dragon-cli/cmd/generate"
	"github.com/svetozar12/dragon-cli/utils"
)

var rootCmd = &cobra.Command{
	Use:   "dragon-cli",
	Short: "Generate full-stack projects based on user choices",
	Run: func(cmd *cobra.Command, args []string) {
		generate.Generate(cmd, args)
	},
}

var Version = "v1.0.2"

func Execute() {
	utils.InitFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
