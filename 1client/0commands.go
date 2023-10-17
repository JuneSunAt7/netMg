package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var ROOT = "filestore/clientDir"

// dynamic root dir
func init() {
	ROOT, _ = filepath.Abs("filestore/clientDir")
}
func Download(conn net.Conn) {

}
func Upload(conn net.Conn) {

}
func ListFile(conn net.Conn) {

}

// HandleClient помощник
func HandleClient(conn net.Conn) {
	stdreader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("ftp> ")
		cmd, _ := stdreader.ReadString('\n')
		cmdArr := strings.Fields(strings.Trim(cmd, "\n"))

		switch strings.ToLower(cmdArr[0]) {

		case "upload":
			if len(cmdArr) != 2 {
				fmt.Println("provide File name please")
				continue
			}
			fmt.Print("\x1b[93m") // F_LightYellow
			fmt.Printf("File Password >> ")
			fmt.Print("\033[0m") // Reset
			fmt.Print("\033[8m") // Hidden
			myFPass, _ := stdreader.ReadString('\n')
			fmt.Print("\033[28m") // Reset_Hidden
			if len(myFPass) < 5 {
				fmt.Println("Too short password")
				continue
			}

			sendFile(conn, cmdArr[1], strings.Trim(myFPass, "\n"))

		case "download":
			if len(cmdArr) != 2 {
				fmt.Println("provide File name please")
				continue
			}

			fmt.Print("\x1b[93m") // F_LightYellow
			fmt.Printf("File Password >> ")
			fmt.Print("\033[0m") // Reset
			fmt.Print("\033[8m") // Hidden
			myFPass, _ := stdreader.ReadString('\n')
			fmt.Print("\033[28m") // Reset_Hidden
			if len(myFPass) < 5 {
				fmt.Println("Too short password")
				continue
			}

			getFile(conn, cmdArr[1], strings.Trim(myFPass, "\n"))

		case "ls":
			conn.Write([]byte(cmd))
			buffer := make([]byte, 4096)
			n, _ := conn.Read(buffer)
			fmt.Print(string(buffer[:n]))

		case "exit", "close":
			conn.Write([]byte("close\n"))
			fmt.Println("Logged out")
			return

		default:
			fmt.Println("Invalid Command, Supported: upload | download | ls | exit")
		}
	}
}
