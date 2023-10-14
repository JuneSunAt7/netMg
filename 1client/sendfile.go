package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func sendFile(conn net.Conn, fname string, myFPass string) {

	content, err := ioutil.ReadFile(ROOT + "/" + fname)
	if err != nil {
		log.Println(err)
		return
	}

	arrEnc, err := CBCEncrypter(myFPass, content)
	if err != nil {
		log.Println(err)
		return
	}

	conn.Write([]byte(fmt.Sprintf("upload %s %d\n", fname, len(arrEnc))))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	str := strings.Trim(string(buf[:n]), "\n")
	commandArr := strings.Fields(str)
	if commandArr[0] != "200" {
		log.Println(str)
		return
	}

	io.Copy(conn, bytes.NewReader(arrEnc))

	n, err = conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(strings.Trim(string(buf[:n]), "\n"))

	checkFileMD5Hash(ROOT + "/" + fname)
}
