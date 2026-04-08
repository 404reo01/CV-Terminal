package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/data"
)

func RenderContact(r *lipgloss.Renderer, width, height int) string {
	title := TitleStyle(r).Render("◈ Contact")
	sep := SeparatorStyle(r).Render(Separator(minC(width-4, 60)))

	var sb strings.Builder
	for _, c := range data.Resume.Contact {
		icon := AccentStyle(r).Render(c.Icon)
		label := SubtitleStyle(r).Render(c.Label + ":")
		value := BodyStyle(r).Render(c.Value)
		sb.WriteString(icon + "  " + label + "  " + value + "\n")
	}

	help := HelpStyle(r).Render("esc  retour")

	page := lipgloss.JoinVertical(lipgloss.Left,
		title,
		sep,
		"",
		sb.String(),
		"",
		help,
	)

	box := r.NewStyle().Padding(1, 3).Render(page)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}

func minC(a, b int) int {
	if a < b {
		return a
	}
	return b
}
