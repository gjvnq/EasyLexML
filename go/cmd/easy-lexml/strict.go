package main

import (
	"io"
	"os"

	"github.com/gjvnq/EasyLexML/go"
	"github.com/spf13/cobra"
)

var strictCmd = &cobra.Command{
	Use:   "strict [input] [output]",
	Short: "Converts an EasyLexML document to its strict/final form.",
	Long:  "Converts an EasyLexML document to its strict/final form.\nIf [output] is not present or is -, stdout will be used.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var input_reader io.Reader

		// Get arguments
		input_path := args[0]
		output_path := ""
		if len(args) > 1 {
			output_path = args[1]
		}
		if output_path == "-" {
			output_path = ""
		}

		// Open input
		if input_path != "-" {
			input_file, err := os.Open(input_path)
			panicIfErr(err)
			defer input_file.Close()
			input_reader = input_file
		} else {
			input_reader = os.Stdin
		}

		// Delete and open output file if needed
		if output_path != "" {
			os.Remove(output_path)

			output_file, err := os.OpenFile(output_path, os.O_RDWR|os.O_CREATE, 0644)
			panicIfErr(err)
			defer output_file.Close()

			// Run
			err = easyLexML.Draft2Strict(input_reader, output_file)
			panicIfErr(err)
		} else {
			// Run
			err = easyLexML.Draft2Strict(input_reader, os.Stdout)
			panicIfErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(strictCmd)
}
