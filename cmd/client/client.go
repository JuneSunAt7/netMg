package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"io"
	"net"
	"os"
	"strings"

	"fmt"

	"atomicgo.dev/keyboard/keys"
	"github.com/fatih/color"
	"github.com/pterm/pterm"

	client "github.com/JuneSunAt7/netMg/1client"
)

func createConfig(ip, port string) {
	conf, err := os.Create("confRD.conf")

	if err != nil {
		color.Red("Ошибка создания файла")
	}

	conf.Write([]byte(port + "\n" + ip))
	defer conf.Close()
	color.Green("Успешное создание файла")
}

func Configure() {
	stdReader := bufio.NewReader(os.Stdin)
	color.Magenta("")
	color.Magenta("")
	color.Magenta("")

	color.Magenta("Текущий файл конфишурации: confRD.conf\nЖелаете изменить его? \n[1 - Изменить 2 - Отмена]")
	cmd, _ := stdReader.ReadString('\n')
	cmdArr := strings.Fields(strings.Trim(cmd, "\n"))

	operation := strings.ToLower(cmdArr[0])

	switch operation {
	case "1":
		stdReader := bufio.NewReader(os.Stdin)
		color.Green("IP   |    PORT ")

		cmd, _ := stdReader.ReadString('\n')
		cmdArr := strings.Fields(strings.Trim(cmd, "\n"))
		ip := strings.ToLower(cmdArr[0])

		port := strings.ToLower(cmdArr[1])

		createConfig(ip, port)
	case "2":
		os.Exit(0)
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

		color.GreenString("TCP TLS Server is Connected @ ", HOST, ":", PORT)
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
			readConfig()
		case "Выход":
			client.Exit(connect)
		}
	}

}

func readConfig() {
	file, err := os.Open("confRD.conf")

	if err != nil {
		color.Red("Ошибка конфигурирования")
		os.Exit(1)
	}

	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		color.Green(string(data[:n]))

	}

}

const (
	PORT = "2121"
	HOST = "localhost"
)

func main() {
	Run()
}
