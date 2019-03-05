package main

import "github.com/spf13/cobra"
import "fmt"

var htmlCmd = &cobra.Command{
	Use:   "html [input] [output]",
	Short: "Converts an EasyLexML document to an HTML page.",
	Long:  "Converts an EasyLexML document to an HTML page.\nIf [output] is not present or is -, stdout will be used.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet.")
	},
}

func init() {
	rootCmd.AddCommand(htmlCmd)
}
