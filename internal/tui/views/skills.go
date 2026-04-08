package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/data"
)

func RenderSkills(r *lipgloss.Renderer, width, height, scrollOffset int) string {
	content := buildSkillsContent(r, width)
	title := TitleStyle(r).Render("◈ Compétences")
	sep := SeparatorStyle(r).Render(Separator(minS(width-4, 60)))

	maxScroll := MaxScroll(content, height-4)
	scrolled := ApplyScroll(content, height-4, scrollOffset)

	scrollHint := ""
	if maxScroll > 0 {
		if scrollOffset < maxScroll {
			scrollHint = HelpStyle(r).Render("▼ suite")
		}
		if scrollOffset > 0 {
			scrollHint = HelpStyle(r).Render("▲ haut  ▼ suite")
			if scrollOffset >= maxScroll {
				scrollHint = HelpStyle(r).Render("▲ haut")
			}
		}
	}

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

func buildSkillsContent(r *lipgloss.Renderer, width int) string {
	var sb strings.Builder
	for _, cat := range data.Resume.Skills {
		sb.WriteString(SubtitleStyle(r).Bold(true).Render("▸ " + cat.Category))
		sb.WriteString("\n")

		tags := []string{}
		for _, item := range cat.Items {
			tags = append(tags, TagStyle(r).Render(item))
		}

		// wrap les tags si trop larges
		line := "  " + strings.Join(tags, " ")
		sb.WriteString(line)
		sb.WriteString("\n\n")
	}
	return strings.TrimRight(sb.String(), "\n")
}

func minS(a, b int) int {
	if a < b {
		return a
	}
	return b
}
