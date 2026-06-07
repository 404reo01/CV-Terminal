package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/data"
)

func RenderProjects(r *lipgloss.Renderer, width, height, scrollOffset int) string {
	content := buildProjectsContent(r, width)
	title := TitleStyle(r).Render("◈ Projets")
	sep := SeparatorStyle(r).Render(Separator(min(width-4, 60)))

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

	_ = fmt.Sprintf
	box := r.NewStyle().Padding(1, 2).Render(page)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}

func buildProjectsContent(r *lipgloss.Renderer, width int) string {
	var sb strings.Builder
	for _, p := range data.Resume.Projects {
		// Nom du projet
		sb.WriteString(TitleStyle(r).Render("▸ " + p.Name))
		sb.WriteString("\n")

		// Description
		sb.WriteString(BodyStyle(r).Render(p.Description))
		sb.WriteString("\n")

		// Tags tech
		tags := []string{}
		for _, t := range p.Tech {
			tags = append(tags, TagStyle(r).Render(t))
		}
		sb.WriteString(strings.Join(tags, " "))
		sb.WriteString("\n")

		if p.Link != "" {
			sb.WriteString(AccentStyle(r).Render("⌥ " + p.Link))
			sb.WriteString("\n")
		}
		if p.Live != "" {
			sb.WriteString(AccentStyle(r).Render("◈ " + p.Live))
			sb.WriteString("\n")
		}

		sb.WriteString(SeparatorStyle(r).Render(Separator(min(width-8, 50))))
		sb.WriteString("\n")
	}
	return sb.String()
}

func MaxScrollProjects(height int) int {
	// approximation basée sur ~6 lignes par projet
	totalLines := len(data.Resume.Projects) * 6
	max := totalLines - (height - 4)
	if max < 0 {
		return 0
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
