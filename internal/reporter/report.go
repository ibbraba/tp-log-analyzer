package reporter

import (
	"encoding/json"
	"os"

	"github.com/ibbraba/tp-log-analyzer/internal/analyzer"
)

func ExportReportToFile(filePath string, report []analyzer.ReportEntry) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return &analyzer.ParsingError{URL: filePath, Err: err}
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return &analyzer.ParsingError{URL: filePath, Err: err}
	}
	return nil
}
