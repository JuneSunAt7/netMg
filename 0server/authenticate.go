package server

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"net"
	"os"

	"github.com/JuneSunAt7/netMg/logger"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredArr []Credentials

func (p *CredArr) FromJSON(r io.Reader) error {
	en := json.NewDecoder(r)
	return en.Decode(p)
}

var Uname string

func GetCred() *CredArr {

	f, _ := os.Open("user_creds.db")
	var creds CredArr
	creds.FromJSON(f)
	return &creds
}

func AuthenticateClient(conn net.Conn) error {

	creds := GetCred()
	if len(*creds) == 0 {
		return errors.New("no credentials: ")
	}

	reader := bufio.NewScanner(conn)

	// Validate user

	reader.Scan()
	uname := reader.Text()

	for _, cred := range *creds {
		Uname = uname

		if CheckUserCert(Uname) {
			logger.Println("Server:Client", uname, "Validated")
			conn.Write([]byte("1"))
			return nil
		} else {
			conn.Write([]byte("0"))
			reader := bufio.NewScanner(conn)

			// Validate user
			// TODO #11 fix validate user3 pass3
			reader.Scan()
			uname := reader.Text()
			reader.Scan()
			passwd := reader.Text()

			if cred.Username == uname && cred.Password == passwd {
				logger.Println("Server:Client", uname, "Validated")
				conn.Write([]byte("1"))
				return nil
			}
		}
	}

	conn.Write([]byte("0"))
	return errors.New("invalid credentials: " + uname)

}
