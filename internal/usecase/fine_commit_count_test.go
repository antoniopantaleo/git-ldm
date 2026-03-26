package usecase

import (
	"errors"
	"testing"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
)

type numberOfCommitsRepoMock struct {
	numberOfCommits domain.FileCommitCount
	err             error
}

func (m *numberOfCommitsRepoMock) FileCommitCount(path string) (domain.FileCommitCount, error) {
	return m.numberOfCommits, m.err
}

func TestFileCommitCountUseCaseHappyPath(t *testing.T) {
	sut := NewFileCommitCountUseCase(
		&numberOfCommitsRepoMock{
			numberOfCommits: 15,
		},
	)
	numberOfCommits, err := sut.Execute("any/path")
	if err != nil {
		t.Fatal(err)
	}
	if numberOfCommits != 15 {
		t.Fatalf("expected number of commits to be 15, got %d", numberOfCommits)
	}
}

func TestFileCommitCountUseCaseEmptyPath(t *testing.T) {
	sut := NewFileCommitCountUseCase(
		&numberOfCommitsRepoMock{},
	)
	_, err := sut.Execute("")
	if err == nil {
		t.Fatal(err)
	}
}

func TestFileCommitCountUseCaseRepoError(t *testing.T) {
	sut := NewFileCommitCountUseCase(
		&numberOfCommitsRepoMock{
			err: errors.New("any error"),
		},
	)
	_, err := sut.Execute("any/path")
	if err == nil {
		t.Fatal(err)
	}
}
