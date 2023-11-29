package client

import (
	"fmt"
	"net"
	"path/filepath"
	"time"

	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

var ROOT = "filestore/clientDir"

// dynamic root dir
func init() {
	ROOT, _ = filepath.Abs("filestore/storeclient")
}

func Upload(conn net.Conn) {
	fname, _ := pterm.DefaultInteractiveTextInput.Show("Имя файла")
	passwd, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Пароль для файла")

	p, _ := pterm.DefaultProgressbar.WithTotal(5).WithTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		if i == 6 {
			time.Sleep(time.Second * 3) // Simulate a slow download.
		}
		p.UpdateTitle("Загрузка в облако")         // Update the title of the progressbar.
		pterm.Success.Println("Загрузка в облако") // If a progressbar is running, each print will be printed above the progressbar.
		p.Increment()                              // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
		time.Sleep(time.Millisecond * 350)         // Sleep 350 milliseconds.
	}
	sendFile(conn, fname, passwd+"\n")

}
func Download(conn net.Conn) {
	fname, _ := pterm.DefaultInteractiveTextInput.Show("Имя файла")
	passwd, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Файловый пароль")
	p, _ := pterm.DefaultProgressbar.WithTotal(5).WithTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		if i == 6 {
			time.Sleep(time.Second * 3) // Simulate a slow download.
		}
		p.UpdateTitle("Выгрузка из облака")         // Update the title of the progressbar.
		pterm.Success.Println("Выгрузка из облака") // If a progressbar is running, each print will be printed above the progressbar.
		p.Increment()                               // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
		time.Sleep(time.Millisecond * 350)          // Sleep 350 milliseconds.
	}
	getFile(conn, fname, passwd+"\n")
}
func ListFiles(conn net.Conn) {
	conn.Write([]byte("ls\n"))
	buffer := make([]byte, 4096)
	n, _ := conn.Read(buffer)
	pterm.FgGreen.Println(string(buffer[:n]))

}

func Exit(conn net.Conn) {
	conn.Write([]byte("close\n"))
	pterm.FgGreen.Println("Выход из системы")

}
func CertPasswd(conn net.Conn) {
	var certoptions []string

	certoptions = append(certoptions, fmt.Sprintf("Доступные сертификаты"))
	certoptions = append(certoptions, fmt.Sprintf("Создать сертификат"))
	certoptions = append(certoptions, fmt.Sprintf("Назад"))

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(certoptions)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(certoptions).Show()
		switch selectedOptions {
		case "Доступные сертификаты":
			AllCertificates(conn)
		case "Создать сертификат":

		case "Назад":
			return
		}
	}

}
func Autoreserved() {
	var options []string

	options = append(options, fmt.Sprintf("Календарь авторезервирования"))
	options = append(options, fmt.Sprintf("Файлы для авторезервирования"))
	options = append(options, fmt.Sprintf("Настройки"))
	options = append(options, fmt.Sprintf("Назад"))

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		switch selectedOptions {
		case "Календарь авторезервирования":

		case "Файлы для авторезервирования":

		case "Настройки":

		case "Назад":
			return
		}
	}

}
