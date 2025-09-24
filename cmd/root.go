package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "config",
	Short: "Crée la configuration du script",
	Long: `Crée la configuration du script en demandant les informations
		à l'utilisateur et en les enregistrant dans un fichier config.json.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
