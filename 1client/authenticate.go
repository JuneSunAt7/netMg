package client

import (
	"errors"
	"net"

	"github.com/pterm/pterm"
)

func getUserCert(conn net.Conn, username string) error {
	netbuff := make([]byte, 1024)
	n, err := conn.Read(netbuff)
	if err != nil {
		return err
	}
	pterm.FgBlue.Println(string(netbuff[:n]))
	conn.Write([]byte(username + "\n"))

	n, err = conn.Read(netbuff)
	if err != nil {
		return err
	}
	if string(netbuff[:n]) == "1" {
		return nil
	} else {
		pterm.FgRed.Println("Сертификат не найден! Используйте пароль пользователя ")
		return errors.New("сертификат не найден")
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
	getUserCert(conn, uname+"\n")
	n, err = conn.Read(buffer)
	if string(buffer[:n]) == "1" {
		pterm.FgGreen.Println("Сертификат найден!")
		return nil
	} else {
		passwd, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Пароль")
		logger := pterm.DefaultLogger
		logger.Info("Выполняется вход", logger.Args("пользователь", uname))
		conn.Write([]byte(uname + "\n" + passwd + "\n"))

		n, err = conn.Read(buffer)
		if err != nil {
			return err
		}

		if string(buffer[:n]) == "1" {
			return nil
		} else {
			pterm.FgRed.Println("Неверный логин или пароль ")
			return errors.New("неверный логин или пароль ")
		}
	}
}
