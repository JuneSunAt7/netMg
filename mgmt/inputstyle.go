package mgmt

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func InputText() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type modelInput struct {
	textInput textinput.Model
	err       error
}

func initialModel() modelInput {
	ti := textinput.New()
	ti.Placeholder = "Настройки"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return modelInput{
		textInput: ti,
		err:       nil,
	}
}

func (m modelInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m modelInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:

			Settings()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m modelInput) View() string {
	return fmt.Sprintf(
		"Введите ваш запрос\n\n%s\n\n%s",
		m.textInput.View(),
		"Нажмите Esc для выхода",
	) + "\n"
}
