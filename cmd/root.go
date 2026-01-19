package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ldm",
		Short: "ldm is a git extension for files history explorations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello world!")
		},
	}
}
