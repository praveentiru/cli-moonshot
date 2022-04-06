package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

var (
	selectedStyle = lipgloss.NewStyle().
		PaddingRight(2).
		Foreground(lipgloss.Color("35")).
		ColorWhitespace(false)
)

func initialModel() model {
	return model{
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {

	lv := m.todoView()
	rv := `This is just a test run to see how two views are joined. 
This is just a test run to see how two views are joined. 
This is just a test run to see how two views are joined. 
This is just a test run to see how two views are joined. 
This is just a test run to see how two views are joined. 
This is just a test run to see how two views are joined. 
This is just a test run to see how two views are joined`
	rv = lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("228")).
		Margin(1, 1, 1, 1).
		Render(rv)

	s := lipgloss.JoinHorizontal(lipgloss.Top, lv, rv)
	return s
}

func (m model) todoView() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		r := fmt.Sprintf("%s [%s] %s", cursor, checked, choice)
		if checked == "x" {
			r = selectedStyle.Render(r)
		}
		s += r
		s += "\n"
	}

	s += "\nPress q to quit.\n"

	s = lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("228")).
		Margin(1, 1, 1, 1).
		Render(s)

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
