package server

import (
	"bufio"
	"net"
	"path/filepath"
	"strings"
	"time"

	"github.com/EwRvp7LV7/45586694crypto/logger"
)

var ROOT = "filestore"

func init() {
	ROOT, _ = filepath.Abs("filestore")
}

func HandleServer(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Server:Connection Established"))

	if err := AuthenticateClient(conn); err != nil {
		logger.Println(err.Error())
		return
	}

	buf := bufio.NewScanner(conn)
	for buf.Scan() {

		commandArr := strings.Fields(strings.Trim(buf.Text(), "\n"))

		conn.SetDeadline(time.Now().Add(time.Minute * 5))

		switch strings.ToLower(commandArr[0]) {

		case "download":
			logger.Println("Download Request")
			sendFile(conn, commandArr[1])

		case "upload":
			logger.Println("Upload Request")
			getFile(conn, commandArr[1], commandArr[2])

		case "ls":
			logger.Println("ls")
			getListFiles(conn)

		case "close":
			logger.Println("closed")
			return
		}
	}
}

// buf := new(bytes.Buffer)

// for {
// 	//in net.Conn is epson methods ReadRunes and like it
// 	//that is why this reading byte to byte. A4.
// 	io.CopyN(buf, conn, 1)
// 	if buf.Bytes()[len(buf.Bytes())-1] == '\n' {

// 		commandArr := strings.Fields(strings.Trim(buf.String(), "\n"))
// 		buf.Reset()

// 		conn.SetDeadline(time.Now().Add(time.Minute * 5))

// 		switch strings.ToLower(commandArr[0]) {

// 		case "download":
// 			logger.Println("Download Request")
// 			sendFile(conn, commandArr[1])

// 		case "upload":
// 			logger.Println("Upload Request")

// 			getFile(conn, commandArr[1], commandArr[2])

// 		case "ls":
// 			logger.Println("ls")
// 			getListFiles(conn)

// 		case "close":
// 			logger.Println("closed")
// 			return
// 		}
// 	}

// }
