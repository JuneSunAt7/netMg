package client

import (
	"errors"
	"net"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
)

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
		color.Red("Неверный логин или пароль ")
		return errors.New("неверный логин или пароль ")
	}
}
