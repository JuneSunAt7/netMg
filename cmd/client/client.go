package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"io"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"

	client "github.com/EwRvp7LV7/45586694crypto/1client"
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
	color.Magenta("Текущий файл конфишурации: confRD.conf\nЖелаете изменить его? \n [1 - Изменить 2 - Отмена]")
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

		color.Green("TCP server is Connected @ ", HOST, ":", PORT)

	} else {

		conf := &tls.Config{
			// InsecureSkipVerify: true,
		}

		connect, err = tls.Dial("tcp", HOST+":"+PORT, conf)
		if err != nil {
			return err
		}

		color.Green("TCP TLS Server is Connected @ ", HOST, ":", PORT)
	}

	defer connect.Close()

	if err := client.AuthenticateClient(connect); err != nil {
		return err
	}

	color.HiCyan("           Доступные функции             ")
	color.Blue("______________________________________")
	color.Blue("|   1    |  Загрузить файл           |")
	color.Blue("|________|___________________________|")
	color.Blue("|   2    |  Скачать файл             |")
	color.Blue("|________|___________________________|")
	color.Blue("|   3    |  Список файлов            |")
	color.Blue("|________|___________________________|")
	color.Blue("|   4    |  Конфигурация сервера     |")
	color.Blue("|________|___________________________|")
	color.Blue("|   5    |  Выход                    |")
	color.Blue("|________|___________________________|")

	for {
		stdReader := bufio.NewReader(os.Stdin)

		color.HiGreen("Номер функции >>> ")

		cmd, _ := stdReader.ReadString('\n')
		cmdArr := strings.Fields(strings.Trim(cmd, "\n"))

		operation := strings.ToLower(cmdArr[0])

		switch operation {
		case "1":
			client.Upload(connect)
		case "2":
			client.Download(connect)
		case "3":
			client.ListFiles(connect)
		case "4":
			Configure()
		case "5":
			client.Exit(connect)
		}
	}

}
func readConfig() []string {
	file, err := os.Open("confRD.conf")
	var ipArray [2]string

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
	return

}

const (
	PORT = "2121"
	HOST = "localhost"
)

func main() {
	Run()
}
