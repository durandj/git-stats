package cli

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/durandj/git-stats/analyzer"
)

// VERSION is the application version string whic his populated at
// build time.
var Version = "UNKNOWN"

var displayVersion bool
var branch string
var reportFormat string
var reportDirectory string

var rootCmd = &cobra.Command{
	Use:   "git-stats",
	Short: "Generates git stats",
	Run: func(cmd *cobra.Command, args []string) {
		if displayVersion {
			fmt.Println(Version)
			return
		}

		reportFormatter, err := analyzer.NewReportFormatter(reportFormat)
		if err != nil {
			fmt.Printf("Unable to get the report formatter:\n%v\n", err)
			os.Exit(1)
		}

		if _, err := os.Stat(reportDirectory); os.IsNotExist(err) {
			err := os.MkdirAll(reportDirectory, os.ModePerm)
			if err != nil {
				fmt.Printf("Unable to create report directory:\n%v\n", err)
			}
		}

		options := analyzer.AnalysisOptions{
			Branch:          branch,
			Path:            ".",
			ReportFormatter: reportFormat,
		}

		stats, err := analyzer.AnalyzeRepository(options)
		if err != nil {
			fmt.Printf("Unable to collect repository stats:\n%v\n", err)
			os.Exit(1)
		}

		err = reportFormatter(reportDirectory, stats)
		if err != nil {
			fmt.Printf("Unable to generate report:\n%v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	flagSet := rootCmd.Flags()

	flagSet.BoolVar(
		&displayVersion,
		"version",
		false,
		"Display the version",
	)

	flagSet.StringVar(
		&branch,
		"branch",
		"master",
		"The branch to run analysis on",
	)

	flagSet.StringVar(
		&reportFormat,
		"report-format",
		analyzer.ReportFormatterJSON,
		"The output type for the report",
	)

	flagSet.StringVar(
		&reportDirectory,
		"report-dir",
		path.Join(".", ".stats"),
		"The directory to save the report to",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
