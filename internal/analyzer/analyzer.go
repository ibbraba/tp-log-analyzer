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
		return CheckResult{InputTarget: target, Status: "FAILED", Err: fmt.Errorf("impossible d'ouvrir le fichier %s: %w", target.Path, err)}
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
		var unreachable *UnreachableURLError
		if errors.As(res.Err, &unreachable) {
			report.Status = "Inaccessible"
			report.ErrorDetails = fmt.Sprintf("Unreachable URL: %v", unreachable.Err)
		} else {
			report.Status = "Error"
			report.ErrorDetails = fmt.Sprintf("Erreur générique: %v", res.Err)
		}
	}

	return report
}
