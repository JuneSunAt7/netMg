package mgmt

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"os/user"
)

// Start protocol
var GetDict = map[uint64]string{
	1: "GETINFO",
	2: "GETSTAT",
	3: "GETFILE",
	4: "GETERROR",
	5: "GETSEC",
	6: "GET2FA",
}

var PostDict = map[int]string{
	11: "POSTINFO",
	22: "POSTDATA",
	33: "COMMENT",
	44: "POST2FA",
	55: "CONFIG",
}
var UserFun = map[string]int{
	"getinfo":     1,
	"stat":        2,
	"getdata":     3,
	"postinfo":    11,
	"postdata":    22,
	"add comment": 33,
}
var secureHash string

func Handshake(address string) {
	nameUsr := Getusr()
	hashname := sha1.New()
	hashname.Write([]byte(nameUsr))
	sha1_hash := hex.EncodeToString(hashname.Sum(nil))

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
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
		panic(err)
	}
	secureHash = string(buff[0:n])

	fmt.Printf("Succesful connect...\n")

}
func Getusr() string {
	user, err := user.Current()
	if err != nil {
		panic(err.Error())
	}
	return user.Username
}

func PostCode(num int, address string) {
	parseCode := PostDict[num]
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for { // cycle for(int i =0; i < 90 second; i++)
		// Send post-data in server
		if n, err := conn.Write([]byte(parseCode)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
	}

}

func GetCode(address string) uint64 {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return 67 // code-signal about error
	}
	defer conn.Close()
	// получем ответ
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		return 67
	}
	data := binary.BigEndian.Uint64(buff[0:n])
	return data
}

// End protocol
