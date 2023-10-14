package client

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

var tests = []string{
	"<html><head></head><body><p>лооООооонг</p></body></html>",
	"По сути, просто []byteс некоторыми методами, чтобы его можно было использовать с другими компонентами.",
	"If you’re sending a streaming response, such as with server-sent events, you’ll need to detect when the client has hung up, and make sure your app server closes the connection promptly. If the server keeps the connection open for 55 seconds without sending any data, you’ll see a request timeout.",
}

var pass = "MyDarkSecret"

func TestCBC(t *testing.T) {

	for _, plaintext := range tests {

		arrEnc, err := CBCEncrypter(pass, []byte(plaintext))
		if err != nil {
			log.Println("EOF", err) //remove
			return
		}
		fmt.Printf("encoded\n%x\n", arrEnc)

		arrDec, err := CBCDecrypter(pass, arrEnc)
		if err != nil {
			log.Println("EOF", err) //remove
			return
		}
		fmt.Printf("decoded\n%s\n", arrDec)

		if plaintext != string(arrDec) {
			t.Error(
				"Еxpected", plaintext,
				"got", string(arrDec),
			)
		}
	}
}

//go test -test.bench=Benchmark
func BenchmarkEncode(b *testing.B) {
	var empty []byte

	for range tests {

		arrEnc, err := CBCEncrypter(pass, []byte("По сути, просто []byteс некоторыми методами, чтобы его можно было использовать с другими компонентами."))
		if err != nil {
			log.Println("EOF", err) //remove
			return
		}
		empty = arrEnc
		// fmt.Printf("encoded\n%x\n", arrEnc)
	}
	//dev/null
	if bytes.EqualFold(empty, []byte{1, 3}) {
		return
	}
}

func BenchmarkDecode(b *testing.B) {
	var empty []byte
	for range tests {
		bt, _ := hex.DecodeString("b58641c6516c67d957e5958314c321c9229df32481d53ad3e72b56cb205af33d22a10b5d6549a8419a3f915e7080ad73b0f3b1b5de79bc150dbe6659314df0a85dc9d542be49e5a012e4e3a32710de47fd8edb087e79f646d75cfca12a8f7e3c7529364ba8022076083d1509ea0b9b6f923d638fb2b76c19cbee168d4bbc5f4d578a5caef1f6697c9bef9977814893ea8ee8133e89fe6e58f836923d2a1bfae456bd37dc8c14748f7270a0a38548e729d60cbc0f98c268fa69b5c425ab87fe8c9d572a693112f7841a39c25e0e949aac")
		arrDec, err := CBCDecrypter(pass, bt)
		if err != nil {
			log.Println("EOF", err) //remove
			return
		}
		empty = arrDec
		fmt.Printf("decoded\n%s\n", arrDec)
	}
	//dev/null
	if bytes.EqualFold(empty, []byte{1, 3}) {
		return
	}
}
