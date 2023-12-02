package client

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

func Calendar() {
	file, err := os.Open(ROOT + "/" + "localSettings" + "/" + "settings.ini")
	if err != nil {
		pterm.FgLightRed.Println("Файл настроек не найден!")
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		pterm.BgCyan.Println("Дни резервирования:")

		pterm.FgGreen.Println(strings.ReplaceAll(string(data[:n]), " ", "\n"))

	}
}
func userfiles() {

	// TODO #5 Add change directory
}
func Setting() {

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
	createSettingsFile(selectedOptions)
}

func createSettingsFile(days []string) {
	/* errmk := os.Mkdir(ROOT+"/"+"localSettings", 0777)
	if errmk != nil {

		fmt.Println(errmk)
		pterm.FgLightRed.Println("Ошибка создания реестра настроек!")
	} */

	outputFile, err := os.Create(ROOT + "/" + "localSettings" + "/" + "settings.ini")
	if err != nil {
		pterm.FgRed.Printfln("Ошибка создания локального файла!")
	}
	defer outputFile.Close()
	outputFile.WriteString(strings.Join(days, " "))

}
