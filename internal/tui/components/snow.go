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

func (s *Snow) Render() string {
	if s.width == 0 || s.height == 0 {
		return ""
	}

	// grille de strings (une cellule = un rune)
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

	rows := make([]string, s.height)
	for y, row := range grid {
		rows[y] = strings.Join(row, "")
	}
	return strings.Join(rows, "\n")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
