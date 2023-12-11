package client

import (
	"errors"
	"hash/fnv"
	"net"
	"regexp"

	"github.com/pterm/pterm"
)

var IsLetter = regexp.MustCompile(`1`).MatchString

func getUserCert(conn net.Conn, username string) bool {
	netbuff := make([]byte, 1024)
	conn.Write([]byte(username + "\n"))

	n, err := conn.Read(netbuff)
	if err != nil {

		return false
	}
	if string(netbuff[:n]) == "1" {

		return true
	} else {
		pterm.FgRed.Println("Сертификат не найден! Используйте пароль пользователя ")

		return false
	}

}

var PASSWD string
var UNAME string

func hash() uint32 {

	h := fnv.New32a()
	h.Write([]byte(UNAME))
	return h.Sum32()
}

func AuthenticateClient(conn net.Conn) error {

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}
	pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgLightGreen)).Println(string(buffer[:n]))

	pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgGreen)).
		WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("Аутенфикация")

	uname, _ := pterm.DefaultInteractiveTextInput.Show("Имя")
	UNAME = uname
	pterm.FgLightBlue.Println("...Поиск сертификата...")
	if getUserCert(conn, uname) {
		pterm.FgGreen.Println("Сертификат найден!")
		return nil
	} else {

		passwd, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Пароль")
		logger := pterm.DefaultLogger
		logger.Info("Выполняется вход", logger.Args("пользователь", uname))
		conn.Write([]byte(passwd + "\n"))

		n, err = conn.Read(buffer)
		if err != nil {
			return err
		}

		if IsLetter(string(buffer[:n])) {
			PASSWD = passwd
			if len(PASSWD) == 0 {
				pterm.FgRed.Println("Ошибка создания криптографического ключа")
				PASSWD = string(hash())
				pterm.FgBlue.Println("Сгенерирован секретный ключ " + PASSWD)
			}

			return nil
		} else {
			pterm.FgRed.Println("Неверный логин или пароль ")
			return errors.New("неверный логин или пароль ")
		}
	}
}
