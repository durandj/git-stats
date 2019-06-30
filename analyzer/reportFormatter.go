package analyzer

import (
	"golang.org/x/xerrors"
)

// ReportFormatter is a function responsible for generating a report
// based on Git stats.
type ReportFormatter = func(reportDirectory string, stats GitStats) error

const (
	// ReportFormatterJSON selects the JSON formatter.
	ReportFormatterJSON string = "json"
)

// NewReportFormatter selects the appropriate formatter.
func NewReportFormatter(formatterType string) (ReportFormatter, error) {
	switch formatterType {
	case ReportFormatterJSON:
		return GenerateJSONReport, nil

	default:
		return nil, xerrors.Errorf(
			"Unrecognized report format type: '%s'",
			formatterType,
		)
	}
}
