package cmd

import (
	"os"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
	"github.com/antoniopantaleo/git-ldm/internal/git"
	"github.com/antoniopantaleo/git-ldm/internal/presenter"
	"github.com/antoniopantaleo/git-ldm/internal/usecase"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "git-ldm",
		Version: "0.1.0-beta",
		Short: "git-ldm is a git extension for files history explorations",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			g, ctx := errgroup.WithContext(ctx)
			cwd, err := os.Getwd()
			filePath := args[0]
			if err != nil {
				return err
			}
			var (
				fc *domain.FileCreation
				fcc domain.FileCommitCount
			)
			repo := git.NewExecGitRepository(&ctx, cwd)
			g.Go(func() error {
				var err error
				usecase := usecase.NewFileCreationUseCase(repo)
				fc, err = usecase.Execute(filePath)
				return err
			})

			g.Go(func() error {
				var err error
				usecase := usecase.NewFileCommitCountUseCase(repo)
				fcc, err = usecase.Execute(filePath)
				return err
			})
			if err := g.Wait(); err != nil {
				return err
			}
			presenter := presenter.NewTermenvPresenter()
			presenter.PresentFileCreation(fc)
			presenter.PresentFileCommitCount(fcc)
			return nil
		},
	}
}
