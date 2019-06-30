package analyzer

import (
	"time"
)

// GitStats is a collection of stats about the Git repository itself.
type GitStats struct {
	Branch       string
	CommitCount  int
	OldestCommit time.Time
	LatestCommit time.Time
}
