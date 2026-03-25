package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
)

type mock struct {
	fileCreation *domain.FileCreation
	err          error
}

func (m *mock) FileCreation(path string) (*domain.FileCreation, error) {
	return m.fileCreation, m.err
}

func TestFileCreationUseCaseHappyPath(t *testing.T) {
	now := time.Now()
	sut := NewFileCreationUseCase(
		&mock{
			fileCreation: &domain.FileCreation{
				Path:      "any/path",
				Author:    "any author",
				CreatedAt: now,
			},
		},
	)
	fileCreation, err := sut.Execute("any/path")
	if err != nil {
		t.Fatal(err)
	}
	if fileCreation.Path != "any/path" {
		t.Fatalf("expected path to be 'any/path', got '%s'", fileCreation.Path)
	}
	if fileCreation.Author != "any author" {
		t.Fatalf("expected author to be 'any author', got '%s'", fileCreation.Author)
	}
	if fileCreation.CreatedAt != now {
		t.Fatalf("expected created at to be '%s', got '%s'", now, fileCreation.CreatedAt)
	}
}

func TestFileCreationUseCaseEmptyPath(t *testing.T) {
	sut := NewFileCreationUseCase(
		&mock{},
	)
	_, err := sut.Execute("")
	if err == nil {
		t.Fatal(err)
	}
}

func TestFileCreationUseCaseRepoError(t *testing.T) {
	sut := NewFileCreationUseCase(
		&mock{
			err: errors.New("any error"),
		},
	)
	_, err := sut.Execute("any/path")
	if err == nil {
		t.Fatal(err)
	}
}
