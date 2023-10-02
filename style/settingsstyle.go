package style

import (
	"fmt"

	"os"

	"github.com/charmbracelet/bubbles/list"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func Settings() {
	items := []list.Item{
		item{title: "Установить сертификат", desc: "Укажите полный путь к сертификату"},
		item{title: "Сконфигурировать с хранилищем", desc: "Необходим сертификат и токен"},
		item{title: "Статистика накопителей", desc: "Сколько занимают места разные типы"},
		item{title: "Статистика использования", desc: "Полная статистика приложения"},
		item{title: "Проверка безопасности хранилища", desc: "Узнайте уровень защиты"},
		item{title: "Авторезервирование", desc: "Настройте частоту резервирования"},
		item{title: "Установить или обновить 2FA", desc: "Установка/обновление биометрии для доступа к данным"},
		item{title: "Проверка технического состояния сервера", desc: "Узнайте основные характеристики сервера"},
		item{title: "Решение проблем с хранилищем", desc: "Исправление ошибок, связанных с сервером"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Настройки"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
