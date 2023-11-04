package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var ROOT = "filestore/clientDir"

// dynamic root dir
func init() {
	ROOT, _ = filepath.Abs("filestore/clientDir")
}
func showMainwindow() {
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
}

func Upload(conn net.Conn) {
	stdReader := bufio.NewReader(os.Stdin)

	color.Magenta("Имя файла >>> ")

	cmd, _ := stdReader.ReadString('\n')
	cmdArr := strings.Fields(strings.Trim(cmd, "\n"))

	filename := strings.ToLower(cmdArr[0])
	color.Magenta("Пароль для файла(не меньше 8 символов) >>> ")
	myFPass, _ := stdReader.ReadString('\n')

	sendFile(conn, filename, strings.Trim(myFPass, "\n"))
	showMainwindow()

}
func Download(conn net.Conn) {
	stdReader := bufio.NewReader(os.Stdin)
	fmt.Printf("\x1b[95m") // fg color magenta
	fmt.Printf("Имя файла >>> ")

	cmd, _ := stdReader.ReadString('\n')
	cmdArr := strings.Fields(strings.Trim(cmd, "\n"))

	filename := strings.ToLower(cmdArr[0])

	fmt.Print("\x1b[94m") // fg color light blue
	fmt.Printf("Файловый пароль  >>> ")
	fmt.Print("\033[0m") // Reset
	fmt.Print("\033[8m") // Hidden
	myFPass, _ := stdReader.ReadString('\n')
	fmt.Print("\x1b[94m") // fg color light blue
	if len(myFPass) < 8 {
		fmt.Println("Слишком короткий пароль")
		return
	}

	getFile(conn, filename, strings.Trim(myFPass, "\n"))
	fmt.Print("\033[0m") // Reset
	showMainwindow()

}
func ListFiles(conn net.Conn) {
	conn.Write([]byte("ls\n"))
	buffer := make([]byte, 4096)
	n, _ := conn.Read(buffer)
	fmt.Print("\x1b[94m")
	fmt.Print(string(buffer[:n]))
	fmt.Print("\033[0m") // Reset
	showMainwindow()
}

func Exit(conn net.Conn) {
	conn.Write([]byte("close\n"))
	fmt.Printf("\x1b[95m") // fg color magenta
	fmt.Println("Выход из системы")
	fmt.Print("\033[0m") // Reset
}
