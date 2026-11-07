package components

import (
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var snowflakes = []string{"❄", "❅", "❆", "·", "∗", "✦"}

type Flake struct {
	x, y       int
	symbol     int
	brightness int // 0=dim, 1=mid, 2=bright
}

type Snow struct {
	flakes   []Flake
	width    int
	height   int
	renderer *lipgloss.Renderer
}

func NewSnow(r *lipgloss.Renderer) *Snow {
	return &Snow{renderer: r}
}

func (s *Snow) Init(width, height int) {
	s.width = width
	s.height = height
	s.flakes = []Flake{}

	count := (width * height) / 80
	for i := 0; i < count; i++ {
		s.flakes = append(s.flakes, Flake{
			x:          rand.Intn(max(width, 1)),
			y:          rand.Intn(max(height, 1)),
			symbol:     rand.Intn(len(snowflakes)),
			brightness: rand.Intn(3),
		})
	}
}

func (s *Snow) Update(width, height int) {
	if width != s.width || height != s.height {
		s.Init(width, height)
		return
	}
	for i := range s.flakes {
		// descente lente
		if rand.Float32() < 0.3 {
			s.flakes[i].y++
			if s.flakes[i].y >= s.height {
				s.flakes[i].y = 0
				s.flakes[i].x = rand.Intn(max(s.width, 1))
			}
		}
		// légère dérive horizontale
		if rand.Float32() < 0.1 {
			drift := rand.Intn(3) - 1
			s.flakes[i].x += drift
			if s.flakes[i].x < 0 {
				s.flakes[i].x = s.width - 1
			}
			if s.flakes[i].x >= s.width {
				s.flakes[i].x = 0
			}
		}
		// variation luminosité
		if rand.Float32() < 0.05 {
			s.flakes[i].brightness = rand.Intn(3)
		}
	}
}

// RenderAround : neige plein écran avec le panneau centré par-dessus.
// La neige tombe autour du panneau, jamais sur son texte.
func (s *Snow) RenderAround(panel string) string {
	lines := strings.Split(panel, "\n")
	panelWidth := lipgloss.Width(panel)
	panelHeight := len(lines)

	// Panneau plus grand que le terminal : pas de place pour la neige
	if panelWidth > s.width || panelHeight > s.height {
		return lipgloss.Place(s.width, s.height, lipgloss.Center, lipgloss.Center, panel)
	}

	// Même centrage que lipgloss.Place
	x0 := (s.width - panelWidth) / 2
	y0 := (s.height - panelHeight) / 2

	grid := s.cells()
	rows := make([]string, s.height)
	for y, row := range grid {
		if y < y0 || y >= y0+panelHeight {
			rows[y] = strings.Join(row, "")
			continue
		}
		line := lines[y-y0]
		line += strings.Repeat(" ", panelWidth-lipgloss.Width(line))
		rows[y] = strings.Join(row[:x0], "") + line + strings.Join(row[x0+panelWidth:], "")
	}
	return strings.Join(rows, "\n")
}

// cells : grille de l'écran, une cellule = un caractère (flocon stylé ou espace).
func (s *Snow) cells() [][]string {
	grid := make([][]string, s.height)
	for y := range grid {
		grid[y] = make([]string, s.width)
		for x := range grid[y] {
			grid[y][x] = " "
		}
	}

	colors := []lipgloss.Color{"#3a5a3a", "#6db87a", "#b0bdb0"}

	for _, f := range s.flakes {
		if f.y < s.height && f.x < s.width {
			symbol := snowflakes[f.symbol]
			color := colors[f.brightness]
			grid[f.y][f.x] = s.renderer.NewStyle().Foreground(color).Render(symbol)
		}
	}
	return grid
}
