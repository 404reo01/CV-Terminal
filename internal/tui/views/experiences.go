package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/data"
)

func RenderExperiences(r *lipgloss.Renderer, width, height, scrollOffset int) string {
	content := buildExperiencesContent(r, width)
	title := TitleStyle(r).Render("◈ Expériences")
	sep := SeparatorStyle(r).Render(Separator(min(width-4, 60)))

	maxScroll := MaxScroll(content, height-4)
	scrolled := ApplyScroll(content, height-4, scrollOffset)
	scrollHint := ScrollHint(r, scrollOffset, maxScroll)

	help := HelpStyle(r).Render("↑/k  ↓/j  défiler   esc  retour")
	footer := lipgloss.JoinHorizontal(lipgloss.Left, help, "   ", scrollHint)

	page := lipgloss.JoinVertical(lipgloss.Left,
		title,
		sep,
		"",
		scrolled,
		"",
		footer,
	)

	box := r.NewStyle().Padding(1, 2).Render(page)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}

func buildExperiencesContent(r *lipgloss.Renderer, width int) string {
	var sb strings.Builder
	for _, e := range data.Resume.Experiences {
		// Entreprise + période
		sb.WriteString(TitleStyle(r).Render("▸ " + e.Company))
		sb.WriteString("  ")
		sb.WriteString(DimStyle(r).Render(e.Period))
		sb.WriteString("\n")

		// Poste
		sb.WriteString(SubtitleStyle(r).Render(e.Role))
		sb.WriteString("\n")

		// Description
		sb.WriteString(BodyStyle(r).Render(e.Description))
		sb.WriteString("\n")

		// Tags tech
		tags := []string{}
		for _, t := range e.Tech {
			tags = append(tags, TagStyle(r).Render(t))
		}
		sb.WriteString(strings.Join(tags, " "))
		sb.WriteString("\n")

		sb.WriteString(SeparatorStyle(r).Render(Separator(min(width-8, 50))))
		sb.WriteString("\n")
	}
	return sb.String()
}
