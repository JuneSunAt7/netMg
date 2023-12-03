package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

func Userfiles() {

	pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).
		WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("Доступные директории:")

	var options []string
	maindir, _ := pterm.DefaultInteractiveTextInput.Show("Директория на устройстве(полный путь до папки или диска)")

	files, err := ioutil.ReadDir(maindir)
	if err != nil {
		pterm.FgRed.Println("Ошибка чтения директорий и файлов!")
	}

	for _, file := range files {
		absPath, err := filepath.Abs(maindir + file.Name())
		if err != nil {
			pterm.FgRed.Println("Ошибка прочтения пути к файлу!")
		}
		options = append(options, fmt.Sprint(absPath+"\n"))

	}

	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()
	pterm.Info.Printfln("Выбранные файлы для резервирования: %s", pterm.Green(selectedOptions))
	updateSettings(selectedOptions)
}

func updateSettings(files []string) {
	// Write path selected files
	outputFile, err := os.Create(ROOT + "/" + "localSettings" + "/" + "path.ini")
	if err != nil {
		pterm.FgRed.Printfln("Ошибка создания локального файла!")
	}
	defer outputFile.Close()
	outputFile.WriteString(strings.Join(files, " "))
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
func Containers() {

}
