package client

import (
	"fmt"

	"github.com/pterm/pterm"
)

func Calendar() {

}
func userfiles() {
	// TODO #5 Add change directory
}
func Setting() {
	// TODO #6 It's include set directories, manage calendar and container files

	pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).
		WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("Дни для резервирования")

	var options []string
	options = append(options, fmt.Sprintf("Понедельник"))
	options = append(options, fmt.Sprintf("Вторник"))
	options = append(options, fmt.Sprintf("Среда"))
	options = append(options, fmt.Sprintf("Четверг"))
	options = append(options, fmt.Sprintf("Пятница"))
	options = append(options, fmt.Sprintf("Суббота"))
	options = append(options, fmt.Sprintf("Воскресенье"))

	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()
	pterm.Info.Printfln("Выбранные дни для резервирования: %s", pterm.Green(selectedOptions))

}
