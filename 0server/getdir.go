package server

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/JuneSunAt7/netMg/logger"
)

func getListFiles(conn net.Conn) {

	files, err := ioutil.ReadDir(ROOT + "/" + Uname) // Read all filenames from filestore.
	if err != nil {
		conn.Write([]byte(err.Error()))
		logger.Println(err.Error())
		return
	}

	fileINFO := ""
	for _, file := range files {
		if !file.IsDir() {
			fileINFO += fmt.Sprintf("%-40s%-25s%-10d\n",
				file.Name(),
				file.ModTime().Format("2006-01-02 15:04:05"),
				file.Size())
		}

	}
	conn.Write([]byte(fileINFO))

}
