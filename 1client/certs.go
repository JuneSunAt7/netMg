package client

import (
	"net"

	"github.com/pterm/pterm"
)

func AllCertificates(conn net.Conn) {
	pterm.FgGreen.Println("Все сертификаты\n")

	conn.Write([]byte("certs\n")) // Bind with server
	buffer := make([]byte, 4096)
	n, _ := conn.Read(buffer)
	pterm.FgGreen.Println(string(buffer[:n])) // Read answer

}
