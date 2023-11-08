package client

import (
	"bufio"
	"errors"
	"net"
	"os"

	"github.com/fatih/color"
)

func AuthenticateClient(conn net.Conn) error {

	stdreader := bufio.NewReader(os.Stdin)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}
	color.Green(string(buffer[:n]))

	color.Green(".....Аутенфикация.....")
	color.HiGreen("Имя >>> ")
	uname, _ := stdreader.ReadString('\n')
	color.HiGreen("Пароль >>> ")
	passwd, _ := stdreader.ReadString('\n')
	conn.Write([]byte(uname))
	conn.Write([]byte(passwd))
	n, err = conn.Read(buffer)
	if err != nil {
		return err
	}

	if string(buffer[:n]) == "1" {
		color.Green("Выполняется вход")
		return nil
	} else {
		color.Red("Неверный логин или пароль " + uname)
		return errors.New("Неверный логин или пароль " + uname)
	}
}
