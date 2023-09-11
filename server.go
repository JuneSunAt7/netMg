package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
	"os/user"
)

var getDict = map[int]string{
	1: "GETINFO",
	2: "GETSTAT",
	3: "GETFILE",
	4: "GETERROR",
}

var postDict = map[int]string{
	1: "POSTINFO",
	2: "POSTDATA",
}

func getusr() string {
	user, err := user.Current()
	if err != nil {
		panic(err.Error())
	}
	return user.Username
}
func geths() {
	nameUsr := getusr()
	hashname := sha1.New()
	hashname.Write([]byte(nameUsr))
	sha1_hash := hex.EncodeToString(hashname.Sum(nil))

	fmt.Println(sha1_hash)

	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		var source = sha1_hash

		// отправляем название компьютера в виде хеша
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		// получем ответ
		fmt.Print("Handshaked hash:")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Print(string(buff[0:n]))
		fmt.Println()
	}
}
func main() {
	geths()

}
