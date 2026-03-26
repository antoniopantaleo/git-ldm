package git

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

var tempPath string

func TestMain(m *testing.M) {
	path, cleanup, err := setup()
	if err != nil {
		panic(err)
	}
	tempPath = path
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func TestRepositoryFileCreation(t *testing.T) {
	sut := NewExecGitRepository(tempPath)
	fileCreation, err := sut.FileCreation("file")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if fileCreation == nil {
		t.Fatalf("Expected file creation, got nil")
	}
	if fileCreation.Path != "file" {
		t.Fatalf("Expected path to be '%s', got '%s'", "file", fileCreation.Path)
	}
	if fileCreation.Author != "any author" {
		t.Fatalf("Expected author to be 'any author', got '%s'", fileCreation.Author)
	}
	if fileCreation.CreatedAt.IsZero() {
		t.Fatalf("Expected created at to be set, got zero value")
	}
}

func TestRepositoryFileCreationNotFound(t *testing.T) {
	sut := NewExecGitRepository(tempPath)
	_, err := sut.FileCreation("non-existent-file")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func setup() (string, func(), error) {
	path, err := os.MkdirTemp("", "git-ldm-test")
	log.Println("Created path in", path)
	cleanup := func() {
		os.RemoveAll(path)
	}
	if err != nil {
		return "", nil, err
	}
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return "", nil, err
	}

	err = os.WriteFile(path+"/file", []byte{}, 0644)
	if err != nil {
		return "", nil, err
	}

	cmd = exec.Command("git", "add", "file")
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return "", nil, err
	}
	cmd = exec.Command("git", "commit", "-m", "any message", "--no-gpg-sign")
	cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=any author", "GIT_AUTHOR_EMAIL=any email")
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return "", nil, err
	}
	return path, cleanup, nil
}
