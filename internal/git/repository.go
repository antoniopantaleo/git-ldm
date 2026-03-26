package git

import (
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
)

type ExecGitRepository struct {
	// The path of the git repository to execute the commands in
	path string
}

func NewExecGitRepository(path string) *ExecGitRepository {
	return &ExecGitRepository{path: path}
}

func (r *ExecGitRepository) FileCreation(path string) (*domain.FileCreation, error) {
	cmd := exec.Command("git", "log", "--diff-filter=A", "--format=%an|%aI", "--", path)
	cmd.Dir = r.path
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if len(output) == 0 {
		return nil, errors.New("file not found")
	}
	outputStr := string(output)
	parts := strings.SplitN(outputStr, "|", 2)
	if len(parts) != 2 {
		return nil, errors.New("unexpected output format")
	}
	author := parts[0]
	date := strings.TrimSpace(parts[1])
	createdAt, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, errors.New("failed to parse created at:")
	}
	fileCreation := &domain.FileCreation{
		Path:      path,
		Author:    author,
		CreatedAt: createdAt,
	}
	return fileCreation, nil
}
