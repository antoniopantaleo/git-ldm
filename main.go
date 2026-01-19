package main

import (
	"os"

	"github.com/antoniopantaleo/git-ldm/cmd"
)

func main() {
	cmd := cmd.NewRootCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
