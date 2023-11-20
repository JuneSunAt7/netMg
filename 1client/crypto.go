package client

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"math"
	"os"
)

func CBCEncrypter(password string, sl []byte) ([]byte, error) {

	key := md5.Sum([]byte(password))

	//требует чтобы размер файла был кратным 16 байтам поэтому в конце дописываем единицу и нули,
	//которые обрежем в Decrypter
	sl16 := make([]byte, int(math.Ceil(float64(len(sl))/aes.BlockSize)*aes.BlockSize)) //%16 bytes
	copy(sl16, sl)
	sl16[len(sl)] = 1

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(sl16)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(sl16))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, (err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], sl16)

	return ciphertext, nil
}

func CBCDecrypter(password string, ciphertext []byte) ([]byte, error) {
	key := md5.Sum([]byte(password))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, (err)
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// without return
	// // CryptBlocks can work in-place if the two arguments are the same.
	// mode.CryptBlocks(ciphertext, ciphertext)
	// //обрезаем см начало CBCEncrypter
	// ciphertext = bytes.TrimRight(ciphertext, "\x00")
	// ciphertext = ciphertext[:len(ciphertext)-1] //именно 1!

	ciphertextOut := make([]byte, len(ciphertext))
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertextOut, ciphertext)
	//обрезаем см начало CBCEncrypter
	ciphertextOut = bytes.TrimRight(ciphertextOut, "\x00")
	ciphertextOut = ciphertextOut[:len(ciphertextOut)-1] //именно 1!
	return ciphertextOut, nil
}

func checkFileMD5Hash(path string) {

	hashFile, _ := os.Open(path)
	defer hashFile.Close()
	h := md5.New()
	if _, err := io.Copy(h, hashFile); err != nil {
		log.Println(err)
	}
	//statsh, _ := hashFile.Stat()
	log.Printf("Хеш: %x\n", h.Sum(nil))
}
