package client

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"log"
	"net"
	"os"
)

func getFile(conn net.Conn, fname string, myFPass string) {

	conn.Write([]byte(fmt.Sprintf("download %s\n", fname)))

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	comStr := strings.Trim(string(buffer[:n]), "\n")
	commandArr := strings.Fields(comStr)

	fileSize, err := strconv.ParseInt(commandArr[2], 10, 64)
	if err != nil || fileSize == -1 {
		log.Println("file size error", err)
		conn.Write([]byte("file size error"))
		return
	}

	conn.Write([]byte("200 Start download!"))

	buf := new(bytes.Buffer)
	io.Copy(buf, io.LimitReader(conn, fileSize))

	arrDec, err := CBCDecrypter(myFPass, buf.Bytes())
	if err != nil {
		log.Println(err)
		return
	}

	outputFile, err := os.Create(ROOT + "/" + fname)
	if err != nil {
		log.Println(err)
	}
	io.Copy(outputFile, bytes.NewReader(arrDec))
	defer outputFile.Close()

	// conn.Write([]byte("File Downloaded successfully"))
	log.Println("File Downloaded successfully")

	checkFileMD5Hash(ROOT + "/" + fname)
}
