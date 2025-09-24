package analyzer

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ibbraba/tp-log-analyzer/internal/config"
)

type CheckResult struct {
	InputTarget config.InputTarget
	Status      string
	Err         error
}

// Utilisé pour l'export
type ReportEntry struct {
	LogId        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func AnalyzeLogFile(target config.InputTarget) CheckResult {
	file, err := os.Open(target.Path)
	if err != nil {
		return CheckResult{InputTarget: target, Status: "FAILED", Err: &FileNotFoundError{URL: target.Path, Err: err}}
	}

	defer file.Close()
	time.Sleep(2 * time.Second) // Simule le temps de traitement
	// Logique d'analyse des logs ici
	fmt.Printf("Analyse du fichier %s terminée avec succès.\n", target.Path)
	return CheckResult{InputTarget: target, Status: "OK", Err: nil}
}

func ConvertToReportEntry(res CheckResult) ReportEntry {
	report := ReportEntry{
		LogId:        res.InputTarget.Id,
		FilePath:     res.InputTarget.Path,
		Status:       res.Status, // Statut par défaut
		Message:      "",
		ErrorDetails: "",
	}

	if res.Err != nil {
		var fileNotFoundError *FileNotFoundError
		var parsingErr *ParsingError
		if errors.As(res.Err, &fileNotFoundError) {

			report.Message = fileNotFoundError.Error()
			report.ErrorDetails = fileNotFoundError.Unwrap().Error()

		} else if errors.As(res.Err, &parsingErr) {

			report.Message = parsingErr.Error()
			report.ErrorDetails = parsingErr.Unwrap().Error()
		} else {

			report.Message = "Unknown error occurred"
			report.ErrorDetails = res.Err.Error()
		}
	}

	return report
}
