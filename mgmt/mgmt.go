package mgmt

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
	"os/user"
)

// Protocol server-client
var secureHash string

var GetDict = map[int]string{
	1: "GETINFO",
	2: "GETSTAT",
	3: "GETFILE",
	4: "GETERROR",
}

var PostDict = map[int]string{
	11: "POSTINFO",
	22: "POSTDATA",
	33: "COMMENT",
}
var UserFun = map[string]int{
	"getinfo":     1,
	"stat":        2,
	"getdata":     3,
	"postinfo":    11,
	"postdata":    22,
	"add comment": 33,
}

func Getusr() string {
	user, err := user.Current()
	if err != nil {
		panic(err.Error())
	}
	return user.Username
}

// End protocol

func Geths() {
	nameUsr := Getusr()
	hashname := sha1.New()
	hashname.Write([]byte(nameUsr))
	sha1_hash := hex.EncodeToString(hashname.Sum(nil))

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
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		secureHash = string(buff[0:n])

		fmt.Printf("Succesful connect...\n")
	}

}

func StateConn(numb int) {
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		source := GetDict[numb]
		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
	}
}
