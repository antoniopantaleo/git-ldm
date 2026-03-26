package git

import (
	"context"
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
)

type ExecGitRepository struct {
	ctx *context.Context
	// The path of the git repository to execute the commands in
	path string
}

func NewExecGitRepository(ctx *context.Context, path string) *ExecGitRepository {
	return &ExecGitRepository{ctx: ctx, path: path}
}

func (r *ExecGitRepository) FileCreation(path string) (*domain.FileCreation, error) {
	cmd := exec.CommandContext(*r.ctx, "git", "log", "--diff-filter=A", "--format=%an|%aI", "--", path)
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

func (r *ExecGitRepository) FileCommitCount(path string) (domain.FileCommitCount, error) {
	cmd := exec.CommandContext(*r.ctx, "git", "rev-list", "--count", "HEAD", "--", path)
	cmd.Dir = r.path
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	res := strings.TrimSpace(string(output))
	count, err := strconv.Atoi(res)
	if err != nil {
		return 0, errors.New("failed to parse commit count")
	}
	return domain.FileCommitCount(count), nil
}