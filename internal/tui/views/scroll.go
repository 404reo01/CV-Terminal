package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func ApplyScroll(content string, height, offset int) string {
	lines := strings.Split(content, "\n")
	total := len(lines)

	if offset < 0 {
		offset = 0
	}
	if offset >= total {
		offset = total - 1
	}

	end := offset + height
	if end > total {
		end = total
	}

	return strings.Join(lines[offset:end], "\n")
}

func MaxScroll(content string, height int) int {
	lines := strings.Split(content, "\n")
	max := len(lines) - height
	if max < 0 {
		return 0
	}
	return max
}

// ScrollHint indique s'il reste du contenu au-dessus et/ou en dessous.
func ScrollHint(r *lipgloss.Renderer, offset, maxScroll int) string {
	switch {
	case maxScroll <= 0:
		return ""
	case offset <= 0:
		return HelpStyle(r).Render("▼ suite")
	case offset >= maxScroll:
		return HelpStyle(r).Render("▲ haut")
	default:
		return HelpStyle(r).Render("▲ haut  ▼ suite")
	}
}
