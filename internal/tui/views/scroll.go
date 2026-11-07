package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Lignes occupées autour du contenu d'une page défilante :
// padding haut/bas (2) + titre + séparateur + ligne vide + ligne vide + pied de page.
const pageChrome = 7

// VisibleLines : nombre de lignes de contenu qui tiennent dans le terminal.
func VisibleLines(height int) int {
	return max(height-pageChrome, 1)
}

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

// renderScrollPage : mise en page commune des écrans défilants (titre, séparateur, contenu, aide).
func renderScrollPage(r *lipgloss.Renderer, width, height int, title, content string, offset int) string {
	visible := VisibleLines(height)
	maxScroll := MaxScroll(content, visible)
	offset = min(offset, maxScroll)

	help := HelpStyle(r).Render("↑/k  ↓/j  défiler   esc  retour")
	footer := lipgloss.JoinHorizontal(lipgloss.Left, help, "   ", ScrollHint(r, offset, maxScroll))

	page := lipgloss.JoinVertical(lipgloss.Left,
		TitleStyle(r).Render(title),
		SeparatorStyle(r).Render(Separator(min(width-4, 60))),
		"",
		ApplyScroll(content, visible, offset),
		"",
		footer,
	)

	box := r.NewStyle().Padding(1, 2).Render(page)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}
