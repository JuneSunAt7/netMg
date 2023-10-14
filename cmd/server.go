package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
	"time"

	server "github.com/EwRvp7LV7/45586694crypto/0server"
	"github.com/EwRvp7LV7/45586694crypto/logger"
)

const (
	PORT = "2121"
)

//tsl
//освоить документацию
func run() (err error) {

	var lstnr net.Listener

	boolTSL := flag.Bool("tls", false, "Set tls")
	flag.Parse()
	if !*boolTSL {

		lstnr, err = net.Listen("tcp", ":"+PORT)
		if err != nil {
			return err
		}

	logger.Println("TCP server is UP @ localhost: " + PORT)

	} else {

		cer, err := tls.LoadX509KeyPair("serts/server.crt", "serts/server.key")
		if err != nil {
			log.Println(err)
			return err
		}

		config := &tls.Config{Certificates: []tls.Certificate{cer}}

		lstnr, err = tls.Listen("tcp", ":"+PORT, config)
		if err != nil {
			return err
		}

	logger.Println("TCP TLS Server is UP @ localhost: " + PORT)

	}

	defer lstnr.Close()

	for {
		//TODO Add limit queue/dispatcher
		connection, err := lstnr.Accept()
		connection.SetDeadline(time.Now().Add(time.Minute * 2))
		if err != nil {
			logger.Println("Client Connection failed")
			continue
		}

		go server.HandleServer(connection)
	}

	// return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}
