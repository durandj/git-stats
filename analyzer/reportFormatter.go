package analyzer

import (
	"golang.org/x/xerrors"
)

type ReportFormatter = func(reportDirectory string, stats GitStats) error

const (
	ReportFormatterJSON string = "json"
)

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
