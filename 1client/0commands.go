package client

import (
	"bufio"

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
	color.Magenta("Имя файла >>> ")

	cmd, _ := stdReader.ReadString('\n')
	cmdArr := strings.Fields(strings.Trim(cmd, "\n"))

	filename := strings.ToLower(cmdArr[0])

	color.Magenta("Файловый пароль  >>> ")
	myFPass, _ := stdReader.ReadString('\n')
	if len(myFPass) < 8 {
		color.Red("Слишком короткий пароль")
		return
	}

	getFile(conn, filename, strings.Trim(myFPass, "\n"))
	showMainwindow()

}
func ListFiles(conn net.Conn) {
	conn.Write([]byte("ls\n"))
	buffer := make([]byte, 4096)
	n, _ := conn.Read(buffer)
	color.Blue(string(buffer[:n]))
	showMainwindow()
}

func Exit(conn net.Conn) {
	conn.Write([]byte("close\n"))
	color.Blue("Выход из системы")
}
