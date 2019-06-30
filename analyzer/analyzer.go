package analyzer

import (
	"time"

	"golang.org/x/xerrors"

	"gopkg.in/src-d/go-git.v4"
	gitObject "gopkg.in/src-d/go-git.v4/plumbing/object"
)

// AnalysisOptions controls how the analysis is performed on the
// repository.
type AnalysisOptions struct {
	Branch          string
	Path            string
	ReportFormatter string
}

// AnalyzeRepository analyzes the repository at the given path and
// collects stats.
func AnalyzeRepository(options AnalysisOptions) (GitStats, error) {
	repository, err := git.PlainOpenWithOptions(
		options.Path,
		&git.PlainOpenOptions{DetectDotGit: true},
	)
	if err != nil {
		return GitStats{}, xerrors.Errorf(
			"Unable to open git repository at %s: %v",
			options.Path,
			err,
		)
	}

	commitIterator, err := repository.Log(&git.LogOptions{Order: git.LogOrderCommitterTime})
	if err != nil {
		return GitStats{}, xerrors.Errorf("Unable to search git history: %v", err)
	}
	defer commitIterator.Close()

	stats := GitStats{
		Branch:       options.Branch,
		CommitCount:  0,
		OldestCommit: time.Now(),
		LatestCommit: time.Time{},
	}

	err = commitIterator.ForEach(func(commit *gitObject.Commit) error {
		stats.CommitCount += 1

		if commit.Committer.When.Before(stats.OldestCommit) {
			stats.OldestCommit = commit.Committer.When
		}

		if commit.Committer.When.After(stats.LatestCommit) {
			stats.LatestCommit = commit.Committer.When
		}

		return nil
	})
	if err != nil {
		return GitStats{}, xerrors.Errorf("Unable to determine git stats: %v", err)
	}

	return stats, nil
}
