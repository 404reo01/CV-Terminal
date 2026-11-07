package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/ascii"
	"github.com/rayan/cv-terminal/internal/data"
)

// Même ordre que menuScreens dans tui/app.go
var menuItems = []string{
	"  Expériences  ",
	"  Projets  ",
	"  Compétences  ",
	"  Contact  ",
}

func RenderHome(r *lipgloss.Renderer, width, height, selected int) string {
	fox := AccentStyle(r).Render(ascii.Fox)

	name := r.NewStyle().
		Foreground(ColorFoxOrange).
		Bold(true).
		Render(data.Resume.Name)

	title := SubtitleStyle(r).Render(data.Resume.Title)
	about := DimStyle(r).Align(lipgloss.Center).Render(data.Resume.About)

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

	// Du plus complet au plus compact : on garde la première version qui tient dans le terminal
	layouts := [][]string{
		{fox, "", name, title, "", about, "", menu, "", help},
		{fox, "", name, title, "", menu, "", help},
		{name, title, "", menu, "", help},
	}
	var content string
	for _, parts := range layouts {
		content = lipgloss.JoinVertical(lipgloss.Center, parts...)
		if lipgloss.Height(content) <= height {
			break
		}
	}

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}
