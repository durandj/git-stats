package analyzer

import (
	"bytes"
	"encoding/json"
	"os"
	"path"

	"golang.org/x/xerrors"
)

func GenerateJSONReport(reportDirectory string, stats GitStats) error {
	reportFileName := path.Join(reportDirectory, "stats.json")
	reportFile, err := os.OpenFile(reportFileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return xerrors.Errorf(
			"Unable to open report file at '%s': %v",
			reportFileName,
			err,
		)
	}
	defer reportFile.Close()

	jsonBytes, err := json.Marshal(stats)
	if err != nil {
		return xerrors.Errorf("Unable to format report: %v", err)
	}

	var buffer bytes.Buffer
	if err := json.Indent(&buffer, jsonBytes, "", "    "); err != nil {
		return xerrors.Errorf("Unable to format report: %v", err)
	}

	if _, err = buffer.WriteTo(reportFile); err != nil {
		return xerrors.Errorf("Unable to write report: %v", err)
	}

	return nil
}
