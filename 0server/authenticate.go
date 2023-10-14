package server

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"net"
	"os"

	"github.com/EwRvp7LV7/45586694crypto/logger"
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

func GetCred() *CredArr {

	f, _ := os.Open("credential.json")
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
	//validate user
	reader.Scan()
	uname := reader.Text()
	reader.Scan()
	passwd := reader.Text()

	for _, cred := range *creds {

		if cred.Username == uname && cred.Password == passwd {
			logger.Println("Server:Client", uname, "Validated")
			conn.Write([]byte("1"))
			return nil
		}
	}

	conn.Write([]byte("0"))
	return errors.New("invalid credentials: " + uname)

}
