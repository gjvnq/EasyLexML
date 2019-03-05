package main

import "github.com/spf13/cobra"
import "fmt"

var txtCmd = &cobra.Command{
	Use:   "txt [input] [output]",
	Short: "Converts an EasyLexML document to a pure text file.",
	Long:  "Converts an EasyLexML document to a pure text file.\nIf [output] is not present or is -, stdout will be used.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet.")
	},
}

func init() {
	rootCmd.AddCommand(txtCmd)
}
