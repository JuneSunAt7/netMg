package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/JuneSunAt7/netMg/logger"
)

func getListCert(conn net.Conn) {
	// This function is similar to the function for getting a list of files.
	files, err := ioutil.ReadDir(CERT + "/" + Uname)
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
func createCert(conn net.Conn) {
	errmk := os.Mkdir(CERT+"/"+Uname, 0777)
	if errmk != nil {
		fmt.Println("Cert was created")
		conn.Write([]byte("0"))

	}
	file, err := os.Create(CERT + "/" + Uname + "/" + "main.crt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())
	conn.Write([]byte("1"))
}
