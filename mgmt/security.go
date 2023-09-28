package mgmt

import (
	"math/rand"
	"os"
	"time"
)

func CreateCert(str string) {
	file, err := os.Create("cert.rd")
	if err != nil {
		os.Exit(1)
	}
	file.WriteString(str)
	defer file.Close()

}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	CreateCert(string(b))

}
