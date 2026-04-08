package views

import (
	"github.com/charmbracelet/lipgloss"
)

// Palette forêt-neige
const (
	ColorForestDeep   = lipgloss.Color("#0d1f0f") // fond sombre forêt
	ColorForestMid    = lipgloss.Color("#1a3320") // vert forêt moyen
	ColorMoss         = lipgloss.Color("#4a7c59") // vert mousse
	ColorLeaf         = lipgloss.Color("#6db87a") // vert feuille clair
	ColorSnowWhite    = lipgloss.Color("#f0f4f0") // blanc neige
	ColorSnowDim      = lipgloss.Color("#b0bdb0") // neige grisée
	ColorBark         = lipgloss.Color("#8b6340") // marron écorce
	ColorBarkLight    = lipgloss.Color("#c49a6c") // marron clair
	ColorFoxOrange    = lipgloss.Color("#e07b39") // orange renard
	ColorFoxLight     = lipgloss.Color("#f0a060") // orange clair renard
	ColorIce          = lipgloss.Color("#c8e6e6") // bleu glace
	ColorNight        = lipgloss.Color("#0a1510") // noir nuit forêt
)

func newStyle(r *lipgloss.Renderer) lipgloss.Style {
	return r.NewStyle()
}

// Styles réutilisables
func TitleStyle(r *lipgloss.Renderer) lipgloss.Style {
	return newStyle(r).Foreground(ColorFoxOrange).Bold(true)
}

func SubtitleStyle(r *lipgloss.Renderer) lipgloss.Style {
	return newStyle(r).Foreground(ColorLeaf).Italic(true)
}

func BodyStyle(r *lipgloss.Renderer) lipgloss.Style {
	return newStyle(r).Foreground(ColorSnowWhite)
}

func DimStyle(r *lipgloss.Renderer) lipgloss.Style {
	return newStyle(r).Foreground(ColorSnowDim)
}

func AccentStyle(r *lipgloss.Renderer) lipgloss.Style {
	return newStyle(r).Foreground(ColorIce)
}

func TagStyle(r *lipgloss.Renderer) lipgloss.Style {
	return newStyle(r).
		Foreground(ColorForestDeep).
		Background(ColorMoss).
		PaddingLeft(1).
		PaddingRight(1)
}

func SeparatorStyle(r *lipgloss.Renderer) lipgloss.Style {
	return newStyle(r).Foreground(ColorMoss)
}

func HelpStyle(r *lipgloss.Renderer) lipgloss.Style {
	return newStyle(r).Foreground(ColorMoss).Italic(true)
}

func Separator(width int) string {
	s := ""
	for i := 0; i < width; i++ {
		s += "─"
	}
	return s
}
