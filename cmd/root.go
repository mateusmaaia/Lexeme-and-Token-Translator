package cmd

import (
	"fmt"
	"github.com/mateusmaaia/Lexeme-and-Token-Translator/cmd/lexicalAnalysis"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lexical-translator [path-to-the-file]",
	Short: "Get the lexical and the token translation",
	Long: `Get the lexical and the token translation`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lexicalAnalysis.Read(args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

