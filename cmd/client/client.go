package main

import (
	"crypto/tls"
	"flag"
	"net"

	"fmt"

	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"

	client "github.com/JuneSunAt7/netMg/1client"
)

const (
	PORT = "2121"
	HOST = "localhost"
)

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
	options = append(options, fmt.Sprintf("Сертификаты и пароли"))
	options = append(options, fmt.Sprintf("Авторезервирование"))
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
			client.Configure()

		case "Сертификаты и пароли":
			client.CertPasswd(connect)
		case "Авторезервирование":
			client.Autoreserved()
		case "Выход":
			client.Exit(connect)
			return
		}
	}

}

func main() {
	Run()
}
