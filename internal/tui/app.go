package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rayan/cv-terminal/internal/tui/components"
	"github.com/rayan/cv-terminal/internal/tui/views"
)

type tickMsg time.Time

const (
	screenHome = iota
	screenExperiences
	screenProjects
	screenSkills
	screenContact
)

// Écran ouvert par chaque entrée du menu (même ordre que menuItems dans views/home.go)
var menuScreens = []int{screenExperiences, screenProjects, screenSkills, screenContact}

type Model struct {
	currentScreen int
	menuIndex     int
	width         int
	height        int
	scrollOffsets map[int]int
	snow          *components.Snow
	renderer      *lipgloss.Renderer
	lastInput     time.Time
	animActive    bool
}

func (m *Model) SetSize(w, h int) {
	m.width = w
	m.height = h
	m.snow.Init(w, h)
}

func NewModel(r *lipgloss.Renderer) Model {
	snow := components.NewSnow(r)
	return Model{
		currentScreen: screenHome,
		menuIndex:     0,
		scrollOffsets: map[int]int{},
		snow:          snow,
		renderer:      r,
		lastInput:     time.Now(),
		animActive:    true,
	}
}

func tick() tea.Cmd {
	return tea.Tick(150*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.snow.Init(msg.Width, msg.Height)
		// Le contenu visible change avec la taille : on recadre les positions de défilement
		for screen, offset := range m.scrollOffsets {
			m.scrollOffsets[screen] = min(offset, m.maxScroll(screen))
		}
		return m, nil

	case tickMsg:
		if m.animActive {
			m.snow.Update(m.width, m.height)
		}
		if time.Since(m.lastInput) > 15*time.Second {
			m.animActive = false
		}
		return m, tick()

	case tea.KeyMsg:
		m.lastInput = time.Now()
		m.animActive = true

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "esc", "b":
			m.currentScreen = screenHome

		case "up", "k":
			if m.currentScreen == screenHome {
				m.menuIndex--
				if m.menuIndex < 0 {
					m.menuIndex = len(menuScreens) - 1
				}
			} else {
				offset := m.scrollOffsets[m.currentScreen]
				if offset > 0 {
					m.scrollOffsets[m.currentScreen] = offset - 1
				}
			}

		case "down", "j":
			if m.currentScreen == screenHome {
				m.menuIndex = (m.menuIndex + 1) % len(menuScreens)
			} else if m.scrollOffsets[m.currentScreen] < m.maxScroll(m.currentScreen) {
				m.scrollOffsets[m.currentScreen]++
			}

		case "enter", " ":
			if m.currentScreen == screenHome {
				m.currentScreen = menuScreens[m.menuIndex]
				m.scrollOffsets[m.currentScreen] = 0
			}
		}
	}

	return m, nil
}

// maxScroll : dernière position de défilement utile pour un écran (0 si tout tient).
func (m Model) maxScroll(screen int) int {
	switch screen {
	case screenExperiences:
		return views.MaxScrollExperiences(m.renderer, m.width, m.height)
	case screenProjects:
		return views.MaxScrollProjects(m.renderer, m.width, m.height)
	case screenSkills:
		return views.MaxScrollSkills(m.renderer, m.height)
	default:
		return 0
	}
}

func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	var panel string
	switch m.currentScreen {
	case screenHome:
		panel = views.RenderHome(m.renderer, m.height, m.menuIndex)
	case screenExperiences:
		panel = views.RenderExperiences(m.renderer, m.width, m.height, m.scrollOffsets[screenExperiences])
	case screenProjects:
		panel = views.RenderProjects(m.renderer, m.width, m.height, m.scrollOffsets[screenProjects])
	case screenSkills:
		panel = views.RenderSkills(m.renderer, m.width, m.height, m.scrollOffsets[screenSkills])
	case screenContact:
		panel = views.RenderContact(m.renderer, m.width)
	}

	return m.snow.RenderAround(panel)
}
