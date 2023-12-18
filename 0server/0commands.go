package server

import (
	"bufio"
	"net"
	"path/filepath"
	"strings"
	"time"

	"github.com/JuneSunAt7/netMg/logger"

	"github.com/pterm/pterm"
)

var ROOT = "filestore"
var CERT = "certificates"

func init() {
	ROOT, _ = filepath.Abs("filestore") // Main directory for users files
	CERT, _ = filepath.Abs("certificates")
}

func HandleServer(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Server:Connection Established"))

	if err := AuthenticateClient(conn); err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	buf := bufio.NewScanner(conn)
	for buf.Scan() {

		commandArr := strings.Fields(strings.Trim(buf.Text(), "\n")) // We receive a client request

		conn.SetDeadline(time.Now().Add(time.Minute * 5))

		switch strings.ToLower(commandArr[0]) {

		case "download":
			pterm.Success.Println("Скачивание из облака")
			sendFile(conn, commandArr[1])

		case "upload":
			pterm.Success.Println("Загрузка в облако")
			getFile(conn, commandArr[1], commandArr[2])
		case "ls":
			pterm.Success.Println("Получение списка файлов")
			getListFiles(conn)
		case "certs":
			logger.Println("Управление сертификатами")
			getListCert(conn)
		case "create":
			logger.Println("Создание сертификата")
			dataCert(conn)
		case "getkey":
			logger.Println("Управление ключами")
			sendKey(conn)
		case "close":
			pterm.Warning.Println("Закрытие соединения")
			return
		}
	}
}
