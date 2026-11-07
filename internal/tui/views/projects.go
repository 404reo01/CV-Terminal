package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/data"
)

func RenderProjects(r *lipgloss.Renderer, width, height, scrollOffset int) string {
	return renderScrollPage(r, width, height, "◈ Projets", buildProjectsContent(r, width), scrollOffset)
}

func MaxScrollProjects(r *lipgloss.Renderer, width, height int) int {
	return MaxScroll(buildProjectsContent(r, width), VisibleLines(height))
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
