package main

import (
	"crypto/tls"
	"flag"
	"io"
	"net"
	"os"

	"fmt"

	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"

	client "github.com/JuneSunAt7/netMg/1client"
)

const (
	PORT = "2121"
	HOST = "localhost"
)

func createConfig(ip, port string) {
	conf, err := os.Create("confRD.conf")

	if err != nil {
		pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgLightRed)).Printfln("Ошибка создания файла")
	}

	conf.Write([]byte(port + "\n" + ip))
	defer conf.Close()
	pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgBlue)).Printfln("Успешное создание файла")
}

func Configure() {
	var options []string

	options = append(options, fmt.Sprintf("Изменить конфигурацию"))
	options = append(options, fmt.Sprintf("Удалить конфигурацию"))
	options = append(options, fmt.Sprintf("Сконфигурировать"))
	options = append(options, fmt.Sprintf("Назад"))

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter

	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		switch selectedOptions {
		case "Изменить конфигурацию":
			ip, _ := pterm.DefaultInteractiveTextInput.Show("IP")
			port, _ := pterm.DefaultInteractiveTextInput.Show("PORT")
			createConfig(ip, port)
			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Println("Конфигурация изменена и сохранена")
		case "Удалить конфигурацию":
			createConfig("", "")
			pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Println("Конфигурация очищена")
		case "Сконфигурировать":
			readConfig()
		case "Назад":
			return
		}
	}
}
func Run() (err error) {

	var connect net.Conn

	boolTSL := flag.Bool("tls", false, "Set tls connection")
	flag.Parse()
	if !*boolTSL {

		connect, err = net.Dial("tcp", HOST+":"+PORT)
		if err != nil {
			return err
		}

	} else {

		conf := &tls.Config{
			// InsecureSkipVerify: true,
		}

		connect, err = tls.Dial("tcp", HOST+":"+PORT, conf)
		if err != nil {
			return err
		}
	}

	defer connect.Close()

	if err := client.AuthenticateClient(connect); err != nil {
		return err
	}

	var options []string

	options = append(options, fmt.Sprintf("Загрузить файл"))
	options = append(options, fmt.Sprintf("Скачать файл"))
	options = append(options, fmt.Sprintf("Список файлов"))
	options = append(options, fmt.Sprintf("Конфигурация"))
	options = append(options, fmt.Sprintf("Выход"))

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter

	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		switch selectedOptions {
		case "Загрузить файл":
			client.Upload(connect)
		case "Скачать файл":
			client.Download(connect)
		case "Список файлов":
			client.ListFiles(connect)
		case "Конфигурация":
			Configure()
		case "Выход":
			client.Exit(connect)
			return
		}
	}

}

func readConfig() {

	file, err := os.Open("confRD.conf")

	if err != nil {
		pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgLightRed)).Println("Ошибка конфигурирования")
		os.Exit(1)
	}

	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Println(n)
		pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen)).Println("Текущая конфигурация:")

		/* arrayConf := strings.Split(string(data[:n]), "\n") */
	}
	/* const (
		PORT = arrayConf[0]
		HOST = arrayConf[1]
	)
	*/
}

func main() {
	Run()
}
