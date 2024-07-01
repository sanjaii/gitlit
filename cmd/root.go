package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitlit",
	Short: "gitlit is a git commit rephrasing tool",
	Long:  `A Fast and git commit rephrasing tool built by sanjaii and picklehari in Go.`,
}

func init() {
	rootCmd.AddCommand(cmdRephrase, cmdSuggest)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
