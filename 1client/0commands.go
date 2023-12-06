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

func init() {
	ROOT, _ = filepath.Abs("filestore/storeclient")
}

func Upload(conn net.Conn) {
	fname, _ := pterm.DefaultInteractiveTextInput.Show("Путь до файла")
	passwd, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Пароль для файла")
	sendFile(conn, fname, passwd+"\n")

	p, _ := pterm.DefaultProgressbar.WithTotal(10).WithTitle("...Загрузка...").Start()

	for i := 0; i < p.Total; i++ {
		if i == 5 {
			time.Sleep(time.Second * 2) // ProgressBar - uploader
		}
		p.UpdateTitle("Загрузка в облако")
		pterm.Success.Println("Загрузка в облако")
		p.Increment()
		time.Sleep(time.Millisecond * 350)
	}
}

func Download(conn net.Conn) {
	fname, _ := pterm.DefaultInteractiveTextInput.Show("Имя файла")
	passwd, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Файловый пароль")
	getFile(conn, fname, passwd+"\n")
	p, _ := pterm.DefaultProgressbar.WithTotal(5).WithTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		if i == 5 {
			time.Sleep(time.Second * 2)
		}
		p.UpdateTitle("Выгрузка из облака") // ProgressBar - downloader
		pterm.Success.Println("Выгрузка из облака")
		p.Increment()
		time.Sleep(time.Millisecond * 350)
	}

}

func ListFiles(conn net.Conn) {
	conn.Write([]byte("ls\n"))
	buffer := make([]byte, 4096)
	n, _ := conn.Read(buffer)
	pterm.FgGreen.Println(string(buffer[:n]))

}

func Exit(conn net.Conn) {
	conn.Write([]byte("close\n"))
	pterm.FgGreen.Println("Выход из облака")
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
			CreateCert(conn)
		case "Назад":
			return
		}
	}
}

func Autoreserved() {
	var options []string

	options = append(options, fmt.Sprintf("Календарь авторезервирования"))
	options = append(options, fmt.Sprintf("Файлы для авторезервирования"))
	options = append(options, fmt.Sprintf("Контейнеры"))
	options = append(options, fmt.Sprintf("Настройки"))
	options = append(options, fmt.Sprintf("Назад"))

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	for {
		selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		switch selectedOptions {
		case "Календарь авторезервирования":
			Calendar()
		case "Файлы для авторезервирования":
			Userfiles()
		case "Контейнеры":
			Containers()
		case "Настройки":
			Setting()
		case "Назад":
			return
		}
	}
}

func CheckDate() {
	// Check date for reserving
}
