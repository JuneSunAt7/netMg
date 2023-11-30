package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/JuneSunAt7/netMg/logger"
)

func sendFile(conn net.Conn, name string) {

	inputFile, err := os.Open(ROOT + "/" + Uname + "/" + name)
	if err != nil {
		logger.Println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	defer inputFile.Close()

	stats, _ := inputFile.Stat()

	conn.Write([]byte(fmt.Sprintf("download %s %d\n", name, stats.Size())))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf) //TODO timeout wating
	if err != nil {
		logger.Println(err.Error())
		return
	}

	str := strings.Trim(string(buf[:n]), "\n")
	commandArr := strings.Fields(str)
	if commandArr[0] != "200" { // Ansver-code 200(succesfully)
		logger.Println(str)
		return
	}

	io.Copy(conn, inputFile)

	logger.Println("File ", name, " Send successfully")
}
