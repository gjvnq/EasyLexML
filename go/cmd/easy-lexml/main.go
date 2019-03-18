package main

import "github.com/spf13/cobra"
import "os"
import "fmt"
import easyLexML "github.com/gjvnq/EasyLexML/go"

var rootCmd = &cobra.Command{
	Use:  os.Args[0],
	Long: "A simple tool for working with EasyLexML documents",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(easyLexML.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&easyLexML.Debug, "debug", "", false, "show debug information")
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
