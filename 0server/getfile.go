package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/EwRvp7LV7/45586694crypto/logger"
)

func getFile(conn net.Conn, name1 string, fs string) {

	fileSize, err := strconv.ParseInt(fs, 10, 64)
	if err != nil || fileSize == -1 {
		logger.Println(err.Error())
		conn.Write([]byte("file size error"))
		return
	}

	name := name1
	fmt.Println(name)

	outputFile, err := os.Create(ROOT + "/" + name)
	if err != nil {
		logger.Println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	defer outputFile.Close()

	conn.Write([]byte("200 Start upload!"))

	//Эта функция использует буфер в 32 КБ
	io.Copy(outputFile, io.LimitReader(conn, fileSize)) //TODO timeout wating

	logger.Println("File " + name + " downloaded successfully")
	// conn.Write([]byte("File Downloaded successfully", fnid))
	fmt.Fprint(conn, "File "+name+" downloaded successfully. UUID file name:\n", name)

}
