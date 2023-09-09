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
		generate.Generate(cmd, args)
	},
}

var verbose bool
var version = "v1.0.2"

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Print the version")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("version") {
			fmt.Printf("Your CLI Tool Version %s\n", version)
			os.Exit(0)
		}
	}
	rootCmd.Flags().String("branch", "master", "Specify the branch name (default is 'master')")

	if err := rootCmd.Execute(); err != nil {

		os.Exit(1)
	}
}
