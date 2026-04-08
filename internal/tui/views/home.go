package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/ascii"
	"github.com/rayan/cv-terminal/internal/data"
)

var menuItems = []string{
	"  Projets  ",
	"  Compétences  ",
	"  Contact  ",
}

const (
	ViewProjects = iota
	ViewSkills
	ViewContact
)

func RenderHome(r *lipgloss.Renderer, width, height, selected int) string {
	// En-tête avec renard
	fox := AccentStyle(r).Render(ascii.Fox)

	name := r.NewStyle().
		Foreground(ColorFoxOrange).
		Bold(true).
		Render(data.Resume.Name)

	title := SubtitleStyle(r).Render(data.Resume.Title)

	header := lipgloss.JoinVertical(lipgloss.Center, fox, "", name, title)

	// Menu
	menuLines := []string{}
	for i, item := range menuItems {
		if i == selected {
			line := r.NewStyle().
				Foreground(ColorNight).
				Background(ColorFoxOrange).
				Bold(true).
				Padding(0, 1).
				Render(item)
			menuLines = append(menuLines, line)
		} else {
			line := r.NewStyle().
				Foreground(ColorSnowWhite).
				Padding(0, 1).
				Render(item)
			menuLines = append(menuLines, line)
		}
	}

	sep := SeparatorStyle(r).Render(Separator(30))
	menu := lipgloss.JoinVertical(lipgloss.Center,
		sep,
		"",
		strings.Join(menuLines, "\n"),
		"",
		sep,
	)

	help := HelpStyle(r).Render("↑/k  ↓/j  naviguer   enter  sélectionner   q  quitter")

	content := lipgloss.JoinVertical(lipgloss.Center,
		header,
		"",
		menu,
		"",
		help,
	)

	_ = fmt.Sprintf // évite import inutilisé
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}
