package domain

type FileCreationRepository interface {
	FileCreation(path string) (*FileCreation, error)
}
type FileCommitCountRepository interface {
	FileCommitCount(path string) (FileCommitCount, error)
}

type Presenter interface {
	PresentFileCreation(fc *FileCreation)
	PresentFileCommitCount(fcc FileCommitCount)
}