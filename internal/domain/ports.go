package domain

type GitRepository interface {
	FileCreation(path string) (*FileCreation, error)
}

type Presneter interface {
	PresentFileCreation(fc *FileCreation)
}