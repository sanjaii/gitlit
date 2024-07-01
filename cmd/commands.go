package cmd

import "github.com/spf13/cobra"

var cmdRephrase = &cobra.Command{
	Use:   "rephrase [commit to rephrase]",
	Short: "Rephrase given commit message",
	Long:  `rephrase is for rephrasing commit based on commit.style rules`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var cmdSuggest = &cobra.Command{
	Use:   "suggest [commit message]",
	Short: "Give suggestions to the given commit message",
	Long:  `Give suggestions to the given commit message based on commit.style rules`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}
