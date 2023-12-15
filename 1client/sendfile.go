package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

func sendFile(conn net.Conn, fname string) {

	// That function use module crypto aka AES & MD5 hasing.
	//The server must make sure that the file is encrypted without errors.
	file := filepath.Base(fname)
	content, err := ioutil.ReadFile(fname)

	if err != nil {
		pterm.Error.Println("Ошибка при загрузке файла")
		fmt.Println(err)
		return
	}

	arrEnc, err := CBCEncrypter(PASSWD, content)
	if err != nil {
		fmt.Println(err)
		pterm.Error.Println("Ошибка при загрузке файла")
		return
	}

	conn.Write([]byte(fmt.Sprintf("upload %s %d\n", file, len(arrEnc))))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)

		pterm.Error.Println("Ошибка при загрузке файла")
		return
	}

	str := strings.Trim(string(buf[:n]), "\n")
	commandArr := strings.Fields(str)
	if commandArr[0] != "200" {
		fmt.Println(err)
		pterm.Error.Println("Ошибка сетевого взаимодействия файла")
		return
	}

	io.Copy(conn, bytes.NewReader(arrEnc))

	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		pterm.Error.Println("Ошибка при загрузке файла")
		return
	}
	p, _ := pterm.DefaultProgressbar.WithTotal(10).WithTitle("...Загрузка...").Start()

	for i := 0; i < p.Total; i++ {
		// Progressbae - uploader
		p.UpdateTitle("Загрузка в облако")
		p.Increment()
		time.Sleep(time.Millisecond * 350)
	}
	pterm.Success.Println(strings.Trim(string(buf[:n]), "\n"))

	checkFileMD5Hash(fname)
}
