package domain

type GitRepository interface {
	FileCreation(path string) (*FileCreation, error)
}
