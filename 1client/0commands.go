package client

import (
	"fmt"
	"net"
	"path/filepath"

	"time"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
)

var ROOT = "filestore/clientDir"

// dynamic root dir
func init() {
	ROOT, _ = filepath.Abs("filestore/clientDir")
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

	fmt.Println(string(buffer[:n]))

	/* pterm.DefaultBarChart.WithBars([]pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}).WithHorizontal().WithWidth(5).Render() */

}

func Exit(conn net.Conn) {
	conn.Write([]byte("close\n"))
	color.Blue("Выход из системы")

}
