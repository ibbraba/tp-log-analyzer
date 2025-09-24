package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ibbraba/tp-log-analyzer/internal/analyzer"
)

func ExportReportToFile(filePath string, report []analyzer.ReportEntry) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("échec de la sérialisation du rapport: %w", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("échec de l'écriture du fichier %s: %w", filePath, err)
	}
	return nil
}
