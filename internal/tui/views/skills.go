package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/data"
)

func RenderSkills(r *lipgloss.Renderer, width, height, scrollOffset int) string {
	return renderScrollPage(r, width, height, "◈ Compétences", buildSkillsContent(r), scrollOffset)
}

func MaxScrollSkills(r *lipgloss.Renderer, height int) int {
	return MaxScroll(buildSkillsContent(r), VisibleLines(height))
}

func buildSkillsContent(r *lipgloss.Renderer) string {
	var sb strings.Builder
	for _, cat := range data.Resume.Skills {
		sb.WriteString(SubtitleStyle(r).Bold(true).Render("▸ " + cat.Category))
		sb.WriteString("\n")

		tags := []string{}
		for _, item := range cat.Items {
			tags = append(tags, TagStyle(r).Render(item))
		}

		sb.WriteString("  " + strings.Join(tags, " "))
		sb.WriteString("\n\n")
	}
	return strings.TrimRight(sb.String(), "\n")
}
