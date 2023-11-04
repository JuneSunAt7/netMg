package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"

	"github.com/fatih/color"

	client "github.com/EwRvp7LV7/45586694crypto/1client"
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

		fmt.Println("TCP server is Connected @ ", HOST, ":", PORT)

	} else {

		conf := &tls.Config{
			// InsecureSkipVerify: true,
		}

		connect, err = tls.Dial("tcp", HOST+":"+PORT, conf)
		if err != nil {
			return err
		}

		fmt.Println("TCP TLS Server is Connected @ ", HOST, ":", PORT)
	}

	defer connect.Close()

	if err := client.AuthenticateClient(connect); err != nil {
		return err
	}
	var operation string
	color.HiBlue("           Доступные функции             ")
	color.Magenta("______________________________________")
	color.Blue("|   1    |  Загрузить файл           |")
	color.Magenta("|________|___________________________|")
	color.Blue("|   2    |  Скачать файл             |")
	color.Magenta("|________|___________________________|")
	color.Blue("|   3    |  Список файлов            |")
	color.Magenta("|________|___________________________|")
	color.Blue("|   4    |  Конфигурация сервера     |")
	color.Magenta("|________|___________________________|")
	color.Blue("|   5    |  Выход                    |")
	color.Magenta("|________|___________________________|")

	for {
		fmt.Scanln(&operation)
		switch operation {
		case "1":
			client.Upload(connect)
		case "2":
			client.Download(connect)
		case "3":
			client.ListFiles(connect)
		case "4":
			client.Exit(connect)
		}
	}

}

const (
	PORT = "2121"
	HOST = "localhost"
)

func main() {
	Run()
}
