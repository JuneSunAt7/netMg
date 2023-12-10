package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/JuneSunAt7/netMg/logger"
)

func getFile(conn net.Conn, name1 string, fs string) {

	fileSize, err := strconv.ParseInt(fs, 10, 64)
	if err != nil || fileSize == -1 { // The size must not be less than zero!
		logger.Println(err.Error())
		conn.Write([]byte("file size error"))
		return
	}

	name := name1
	fmt.Println(name)

	errmk := os.Mkdir(ROOT+"/"+Uname, 0777)
	if errmk != nil {
		fmt.Println("Error create dir")
	}

	outputFile, err := os.Create(ROOT + "/" + Uname + "/" + name)
	if err != nil {
		logger.Println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	defer outputFile.Close()

	conn.Write([]byte("200 Start upload!"))

	//Use buff size 32 bytes
	io.Copy(outputFile, io.LimitReader(conn, fileSize))

	logger.Println("Файл  " + name + " загружен в облако")
	fmt.Fprint(conn, "Файл  "+name+" загружен в облако успешно\n")

}
func getFileWithCert(conn net.Conn, name1 string) {
	sendKey(conn)

	fileSize, err := strconv.ParseInt(fs, 10, 64)
	if err != nil || fileSize == -1 { // The size must not be less than zero!
		logger.Println(err.Error())
		conn.Write([]byte("file size error"))
		return
	}

	name := name1
	fmt.Println(name)

	errmk := os.Mkdir(ROOT+"/"+Uname, 0777)
	if errmk != nil {
		fmt.Println("Error create dir")
	}

	outputFile, err := os.Create(ROOT + "/" + Uname + "/" + name)
	if err != nil {
		logger.Println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	defer outputFile.Close()

	conn.Write([]byte("200 Start upload!"))

	//Use buff size 32 bytes
	io.Copy(outputFile, io.LimitReader(conn, fileSize))

	logger.Println("Файл  " + name + " загружен в облако")
	fmt.Fprint(conn, "Файл  "+name+" загружен в облако успешно\n")

}
