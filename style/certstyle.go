package style

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func cerertTable() {
	columns := []table.Column{
		{Title: "№", Width: 8},
		{Title: "Команда", Width: 30},
	}

	rows := []table.Row{
		{"1", "Генерация сертификата"},
		{"2", "Генерация токена"},
		{"3", "Указать сертификат"},
		{"4", "Назад"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(3),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Ошибка запуска ", err)
		os.Exit(1)
	}
}
func CreateCertTable() {
	cerertTable()
}
