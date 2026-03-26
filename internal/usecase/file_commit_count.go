package usecase

import (
	"errors"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
)

type FileCommitCountUseCase struct {
	repo domain.FileCommitCountRepository
}

func NewFileCommitCountUseCase(repo domain.FileCommitCountRepository) *FileCommitCountUseCase {
	return &FileCommitCountUseCase{repo: repo}
}

func (u *FileCommitCountUseCase) Execute(path string) (domain.FileCommitCount, error) {
	if path == "" {
		return 0, errors.New("Path can not be empty")
	}
	return u.repo.FileCommitCount(path)
}
