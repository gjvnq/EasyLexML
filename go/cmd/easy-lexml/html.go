package main

import (
	"os"

	"github.com/gjvnq/EasyLexML/go"
	"github.com/spf13/cobra"
)

var htmlCmd = &cobra.Command{
	Use:   "html [input] [output]",
	Short: "Converts an EasyLexML document to an HTML page.",
	Long:  "Converts an EasyLexML document to an HTML page.\nIf [output] is not present or is -, stdout will be used.",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

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
		input_file, err := os.Open(input_path)
		panicIfErr(err)
		defer input_file.Close()

		// Delete and open output file if needed
		if output_path != "" {
			os.Remove(output_path)

			output_file, err := os.OpenFile(output_path, os.O_RDWR|os.O_CREATE, 0644)
			panicIfErr(err)
			defer output_file.Close()

			// Run
			err = easyLexML.Strict2HTML(input_file, output_file)
			panicIfErr(err)
		} else {
			// Run
			err = easyLexML.Strict2HTML(input_file, os.Stdout)
			panicIfErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(htmlCmd)
}
