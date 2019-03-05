package main

import "github.com/spf13/cobra"
import "fmt"

var strictCmd = &cobra.Command{
	Use:   "strict [input] [output]",
	Short: "Converts an EasyLexML document to its strict/final form.",
	Long:  "Converts an EasyLexML document to its strict/final form.\nIf [output] is not present or is -, stdout will be used.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet.")
	},
}

func init() {
	rootCmd.AddCommand(strictCmd)
}
