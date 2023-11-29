package server

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/JuneSunAt7/netMg/logger"
)

func getListCert(conn net.Conn) {

	files, err := ioutil.ReadDir(CERT)
	if err != nil {
		conn.Write([]byte(err.Error()))
		logger.Println(err.Error())
		return
	}

	fileINFO := ""
	for _, file := range files {
		if !file.IsDir() {
			fileINFO += fmt.Sprintf("%-40s%-25s\n",
				file.Name(),
				file.ModTime().Format("2006-01-02 15:04:05"))
		}

	}
	conn.Write([]byte(fileINFO))

}
func changePass(conn net.Conn) {
	reader := bufio.NewScanner(conn)
	//validate user
	reader.Scan()
	uname := reader.Text()
	logger.Println(uname)
	reader.Scan()
	passwd := reader.Text()
	logger.Println(passwd)
}
