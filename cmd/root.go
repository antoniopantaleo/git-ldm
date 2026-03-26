package cmd

import (
	"os"
	"sync"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
	"github.com/antoniopantaleo/git-ldm/internal/git"
	"github.com/antoniopantaleo/git-ldm/internal/presenter"
	"github.com/antoniopantaleo/git-ldm/internal/usecase"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "git-ldm",
		Version: "0.1.0-beta",
		Short: "git-ldm is a git extension for files history explorations",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			filePath := args[0]
			if err != nil {
				return err
			}
			var wg sync.WaitGroup
			wg.Add(2)
			repo := git.NewExecGitRepository(cwd)
			var (
				fc *domain.FileCreation
				fcErr error
				fcc domain.FileCommitCount
				fccErr error
			)
			go func() {
				defer wg.Done()
				usecase := usecase.NewFileCreationUseCase(repo)
				fc, fcErr = usecase.Execute(filePath)
				
			}()
			go func() {
				defer wg.Done()
				usecase := usecase.NewFileCommitCountUseCase(repo)
				fcc, fccErr = usecase.Execute(filePath)
				
			}()
			wg.Wait()
			if fcErr != nil {
				return fcErr
			}
			if fccErr != nil {
				return fccErr
			}
			presenter := presenter.NewTermenvPresenter()
			presenter.PresentFileCreation(fc)
			presenter.PresentFileCommitCount(fcc)
			return nil
		},
	}
}
