package presenter

import (
	"fmt"
	"os"

	"github.com/antoniopantaleo/git-ldm/internal/domain"
	"github.com/muesli/termenv"
)

type TermenvPresenter struct {
}

func NewTermenvPresenter() *TermenvPresenter {
	return &TermenvPresenter{}
}

func (p *TermenvPresenter) PresentFileCreation(fc *domain.FileCreation) {
	o := termenv.NewOutput(os.Stdout)

	label := o.String("📄 File").Bold()
	path := o.String(fc.Path).Foreground(o.Color("#89B4FA"))

	authorLabel := o.String("👤 Author").Bold()
	author := o.String(fc.Author).Foreground(o.Color("#A6E3A1"))

	dateLabel := o.String("🗓️  Created").Bold()
	date := o.String(fc.CreatedAt.Format("02 Jan 2006, 15:04")).Foreground(o.Color("#F9E2AF"))

	fmt.Printf("%s    %s\n", label, path)
	fmt.Printf("%s  %s\n", authorLabel, author)
	fmt.Printf("%s %s\n", dateLabel, date)
}

func (p *TermenvPresenter) PresentFileCommitCount(fcc domain.FileCommitCount) {
	o := termenv.NewOutput(os.Stdout)

	label := o.String("🔢 Number of commits").Bold()
	count := o.String(fmt.Sprintf("%d", fcc)).Foreground(o.Color("#F38BA8"))

	fmt.Printf("%s: %s\n", label, count)
}