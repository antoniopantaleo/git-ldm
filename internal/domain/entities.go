package domain

import "time"

type FileCreation struct {
	Path      string
	Author    string
	CreatedAt time.Time
}

type FileCommitCount int