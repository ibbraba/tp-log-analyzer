package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les fichiers de log",
	Long:  `Analyse les fichiers de log et les affiche de manière lisible.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		if path == "" {
			cmd.Help()
			return
		}

		fmt.Printf("Analyse du fichier de log : %s\n", path)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringP("path", "p", "", "Fichier de log à analyser")
}
