package client

import (
	"errors"
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
			return nil
		} else {
			pterm.FgRed.Println("Неверный логин или пароль ")
			return errors.New("неверный логин или пароль ")
		}
	}
}
