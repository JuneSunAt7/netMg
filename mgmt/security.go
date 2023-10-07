package mgmt

import (
	b64 "encoding/base64"
	"math/rand"
	"os"
	"time"
)

const lenRunes = 40

func createCert(str string) {
	file, err := os.Create("cert.rd")
	if err != nil {
		os.Exit(1)
	}
	file.WriteString(str)
	defer file.Close()

}

func LogHashChains() {
	if secureHash == "" {
		return
	}
	file, err := os.Create("hash_chains.rd")
	if err != nil {
		os.Exit(1)
	}
	data := secureHash + " " + time.Now().Format(time.DateTime) + "\n"
	file.WriteString(data)
	defer file.Close()

}
func ChainsUpdate() {
	if _, err := os.Stat("hash_chains.rd"); err != nil {
		if os.IsNotExist(err) {
			RandStringRunes()
		} else {
			upd, err := os.OpenFile("hash_chains.rd", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			defer upd.Close()

			if _, err = upd.WriteString(secureHash + " " + time.Now().Format(time.DateTime) + "\n"); err != nil {
				panic(err)
			}
		}
	}
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes() {
	b := make([]rune, lenRunes)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	createCert(string(b))

}

func genToken() string {
	data, err := os.ReadFile("cert.rd")
	if err != nil {
		os.Exit(1)
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc

}

func UpdateCert(filename string) { // Full path of certificate
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			RandStringRunes()
		} else {
			upd, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			defer upd.Close()

			if _, err = upd.WriteString(genToken()); err != nil {
				panic(err)
			}
		}
	}
}
