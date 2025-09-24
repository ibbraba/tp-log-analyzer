package cmd

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ibbraba/tp-log-analyzer/internal/analyzer"
	"github.com/ibbraba/tp-log-analyzer/internal/config"
	"github.com/ibbraba/tp-log-analyzer/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	inputFilePath  string
	outputFilePath string
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

		targets, err := config.LoadTargetsFromFile(inputFilePath)
		if err != nil {
			fmt.Printf("Erreur lors du chargement des cibles : %v\n", err)
			return
		}

		//Creer waitgroup
		var wg sync.WaitGroup
		resultsChan := make(chan analyzer.CheckResult, len(targets))
		wg.Add(len(targets))

		for _, target := range targets {
			go func(t config.InputTarget) {
				defer wg.Done()
				result := analyzer.AnalyzeLogFile(target)
				resultsChan <- result
			}(target)
		}

		wg.Wait()
		close(resultsChan)

		var finalReport []analyzer.ReportEntry

		for result := range resultsChan {
			reportEntry := analyzer.ConvertToReportEntry(result)
			finalReport = append(finalReport, reportEntry)
			if result.Err != nil {
				var fileNotFoundError *analyzer.FileNotFoundError
				var parsingError *analyzer.ParsingError
				if errors.As(result.Err, &fileNotFoundError) {
					reportEntry.Status = "File Not Found"
					reportEntry.Message = "Le fichier est introuvable."
				} else if errors.As(result.Err, &parsingError) {
					reportEntry.Status = "Parsing Error"
					reportEntry.Message = "Erreur lors de l'analyse du fichier."
				} else {
					reportEntry.Status = "Unknown Error"
					reportEntry.Message = "Une erreur inconnue est survenue."
				}

				fmt.Printf("Erreur lors de l'analyse du fichier %s: %v\n", result.InputTarget.Path, result.Err)
			} else {
				fmt.Printf("Analyse du fichier %s terminée avec succès.\n", result.InputTarget.Path)
			}
		}

		// Exporter le rapport final
		if outputFilePath != "" {
			err = reporter.ExportReportToFile(outputFilePath, finalReport)
			if err != nil {
				fmt.Printf("Erreur lors de l'exportation du rapport : %v\n", err)
			} else {
				fmt.Printf("Rapport exporté avec succès vers %s\n", outputFilePath)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&inputFilePath, "path", "p", "", "Fichier de log à analyser")
	analyzeCmd.Flags().StringVarP(&outputFilePath, "output", "o", "report.json", "Fichier de sortie pour le rapport d'analyse")
	analyzeCmd.MarkFlagRequired("path")
}
