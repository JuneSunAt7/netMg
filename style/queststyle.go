package style

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#a200ff")).
	Faint(true).Align(lipgloss.Center)

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.table.SelectedRow()[0] == "1" {

				CreateCertTable()
			}

		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func CreateTable() {
	columns := []table.Column{
		{Title: "№", Width: 8},
		{Title: "Команда", Width: 20},
		{Title: "Описание", Width: 40},
	}

	rows := []table.Row{
		{"1", "Сертификаты", "Безопасность и программная защита"},
		{"2", "2FA", "Управление паролями и RDevice"},
		{"3", "Конфигурация", "Настройка конфигурации"},
		{"4", "Авторезервирование", "Частота резервирования"},
		{"5", "Статистика", "Статистика накопителей"},
		{"6", "Данные", "Упаковка данных"},
		{"7", "Помощь", "Устранение типичных проблем"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
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
