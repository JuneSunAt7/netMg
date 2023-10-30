package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var tableStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

type modelTable struct {
	table table.Model
}

func (m modelTable) Init() tea.Cmd { return nil }

func (m modelTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			switch m.table.SelectedRow()[0] {

			case "#1":

			case "#2":

			case "#3":

			case "#4":

			}

		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m modelTable) View() string {
	return tableStyle.Render(m.table.View()) + "\n"
}

func MainWindow() {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Команда", Width: 15},
		{Title: "Описание", Width: 30},
	}

	rows := []table.Row{
		{"#1", "Скачать", "Скачивание файла с хранилища"},
		{"#2", "Загрузить", "Загрузка в хранилище"},
		{"#3", "Аутефикация", "Пароли и 2FA"},
		{"#4", "Список", "Список файлов на сервере"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(4),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("#4103fc")).
		Bold(false)
	t.SetStyles(s)

	m := modelTable{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
