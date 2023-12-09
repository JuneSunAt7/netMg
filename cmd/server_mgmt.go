package main

import (
	"fmt"

	"atomicgo.dev/keyboard/keys"
	client "github.com/JuneSunAt7/netMg/1client"
	"github.com/pterm/pterm"
)

func TuiCommands() {
	var options []string

	options = append(options, fmt.Sprintf("Добавить пользователя"))
	options = append(options, fmt.Sprintf("Удалить пользователя"))
	options = append(options, fmt.Sprintf("Данные о пользователях"))
	options = append(options, fmt.Sprintf("Конфигурация"))
	options = append(options, fmt.Sprintf("Выход"))

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.TextStyle.Add(*pterm.NewStyle(pterm.FgBlue))
	printer.KeyConfirm = keys.Enter

	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

		switch selectedOptions {
		case "Добавить пользователя":
			client.Upload(connect)
		case "Удалить пользователя":
			client.Download(connect)
		case "Данные о пользователях":
			client.ListFiles(connect)
		case "Конфигурация":
			client.Configure()
		case "Выход":
			return
		}
	}

}
