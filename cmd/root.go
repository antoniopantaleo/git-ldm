package cmd

import (
	"os"

	"github.com/antoniopantaleo/git-ldm/internal/git"
	"github.com/antoniopantaleo/git-ldm/internal/presenter"
	"github.com/antoniopantaleo/git-ldm/internal/usecase"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ldm",
		Short: "ldm is a git extension for files history explorations",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			filePath := args[0]
			if err != nil {
				return err
			}
			repo := git.NewExecGitRepository(cwd)
			usecase := usecase.NewFileCreationUseCase(repo)
			fileCreation, err := usecase.Execute(filePath)
			if err != nil {
				return err
			}
			presenter := presenter.NewTermenvPresenter()
			presenter.PresentFileCreation(fileCreation)
			return nil
		},
	}
}
