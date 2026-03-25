package usecase

import (
	"errors"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
)

type FileCreationUseCase struct {
	repo domain.GitRepository
}

func NewFileCreationUseCase(repo domain.GitRepository) *FileCreationUseCase {
	return &FileCreationUseCase{repo: repo}
}

func (u *FileCreationUseCase) Execute(path string) (*domain.FileCreation, error) {
	if path == "" {
		return nil, errors.New("path can not be empty")
	}
	return u.repo.FileCreation(path)
}
