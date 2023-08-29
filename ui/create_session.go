package ui

import (
	"fmt"
	"log"

	"cgpa-calc/database"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type CreateSessionModel struct {
	textInput textinput.Model
	err       error
	c         database.CgpaRepo
}

func InitialCSModel(c database.CgpaRepo) CreateSessionModel {
	ti := textinput.New()
	ti.Placeholder = "First 2021/22"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return CreateSessionModel{
		textInput: ti,
		err:       nil,
		c:         c,
	}
}

func (m CreateSessionModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CreateSessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			if msg.Type == tea.KeyEnter {
				m.c.AddSemester(database.NewSemester(m.textInput.Value()))
				log.Printf("you entered %s", m.textInput.Value())
			}
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m CreateSessionModel) View() string {
	return fmt.Sprintf(
		"Enter the session name\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
