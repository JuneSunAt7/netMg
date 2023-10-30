package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"

	client "github.com/EwRvp7LV7/45586694crypto/1client"
)

func Run() (err error) {

	var connect net.Conn

	boolTSL := flag.Bool("tls", false, "Set tls connection")
	flag.Parse()
	if !*boolTSL {

		connect, err = net.Dial("tcp", HOST+":"+PORT)
		if err != nil {
			return err
		}

		fmt.Println("TCP server is Connected @ ", HOST, ":", PORT)

	} else {

		conf := &tls.Config{
			// InsecureSkipVerify: true,
		}

		connect, err = tls.Dial("tcp", HOST+":"+PORT, conf)
		if err != nil {
			return err
		}

		fmt.Println("TCP TLS Server is Connected @ ", HOST, ":", PORT)
	}

	defer connect.Close()

	if err := client.AuthenticateClient(connect); err != nil {
		return err
	}

	client.Upload(connect)

	return nil
}

const (
	PORT = "2121"
	HOST = "localhost"
)

func main() {
	// flag.Parse()

}
