package client

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	// "path/filepath"
)

//AuthenticateClient помощник аутенификации
func AuthenticateClient(conn net.Conn) error {

	stdreader := bufio.NewReader(os.Stdin)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	fmt.Println(string(buffer[:n]))
	fmt.Println("Authentication Required")

	fmt.Print("\x1b[93m") // F_LightYellow
	fmt.Printf("Username >> ")
	fmt.Print("\033[0m") // Reset
	uname, _ := stdreader.ReadString('\n')
	fmt.Print("\x1b[93m") // F_LightYellow
	fmt.Printf("Password >> ")
	fmt.Print("\033[0m") // Reset
	fmt.Print("\033[8m") // Hidden
	passwd, _ := stdreader.ReadString('\n')
	fmt.Print("\033[28m") // Reset_Hidden
	conn.Write([]byte(uname))
	conn.Write([]byte(passwd))
	n, err = conn.Read(buffer)
	if err != nil {
		return err
	}

	if string(buffer[:n]) == "1" {
		fmt.Println("Authentication Successful")
		return nil
	} else {
		// fmt.Println("Invalid credentials. Bye!")
		return errors.New("invalid credentials " + uname)
	}
}
