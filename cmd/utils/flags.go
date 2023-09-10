package utils

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

func InitFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Print the version")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("version") {
			fmt.Printf("Your CLI Tool Version %s\n", cmd.Version)
			os.Exit(0)
		}
	}
	rootCmd.Flags().String("branch", "master", "Specify the branch name (default is 'master')")

	// projectName
	rootCmd.Flags().String("projectName", "", "Specify the project name")
	// beFramework
	rootCmd.Flags().String("beFramework", "", "Specify the backend framework")
	// feFramework
	rootCmd.Flags().String("feFramework", "", "Specify the frontend framework")
	// installDeps
	rootCmd.Flags().String("installDeps", "", "If you want to install your dependencies")
}
