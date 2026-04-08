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
	screenProjects
	screenSkills
	screenContact
)

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
					m.menuIndex = 2
				}
			} else {
				offset := m.scrollOffsets[m.currentScreen]
				if offset > 0 {
					m.scrollOffsets[m.currentScreen] = offset - 1
				}
			}

		case "down", "j":
			if m.currentScreen == screenHome {
				m.menuIndex = (m.menuIndex + 1) % 3
			} else {
				m.scrollOffsets[m.currentScreen]++
			}

		case "enter", " ":
			if m.currentScreen == screenHome {
				switch m.menuIndex {
				case 0:
					m.currentScreen = screenProjects
					m.scrollOffsets[screenProjects] = 0
				case 1:
					m.currentScreen = screenSkills
					m.scrollOffsets[screenSkills] = 0
				case 2:
					m.currentScreen = screenContact
				}
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	bg := m.snow.Render()

	var overlay string
	switch m.currentScreen {
	case screenHome:
		overlay = views.RenderHome(m.renderer, m.width, m.height, m.menuIndex)
	case screenProjects:
		overlay = views.RenderProjects(m.renderer, m.width, m.height, m.scrollOffsets[screenProjects])
	case screenSkills:
		overlay = views.RenderSkills(m.renderer, m.width, m.height, m.scrollOffsets[screenSkills])
	case screenContact:
		overlay = views.RenderContact(m.renderer, m.width, m.height)
	}

	_ = bg
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
		overlay,
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(views.ColorForestDeep),
	)
}
