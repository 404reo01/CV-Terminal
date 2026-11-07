package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/data"
)

func RenderContact(r *lipgloss.Renderer, width int) string {
	title := TitleStyle(r).Render("◈ Contact")
	sep := SeparatorStyle(r).Render(Separator(min(width-4, 60)))

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

	return r.NewStyle().Padding(1, 3).Render(page)
}
