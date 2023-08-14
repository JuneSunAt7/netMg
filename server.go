package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	key := []byte("0123456789abcdef0123456789abcdef") // 256-bit (32 byte) AES ключ
	plaintext := []byte("texttexttexttext")           // исходный текст для шифрования

	// Шифрование
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		fmt.Println("Ошибка шифрования:", err)
		return
	}

	fmt.Printf("Зашифрованный текст в шестнадцатеричном формате: %s\n", hex.EncodeToString(ciphertext))

	// Дешифрование
	decryptedText, err := decrypt(key, ciphertext)
	if err != nil {
		fmt.Println("Ошибка дешифрования:", err)
		return
	}

	fmt.Println("Расшифрованный текст:", string(decryptedText))
}

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("зашифрованный текст слишком короткий")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}
