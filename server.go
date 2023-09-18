package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func getClient(config *oauth2.Config) *http.Client {
	tok, err := tokenFromFile("token.json") // Загрузите токен из файла (если используется)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken("token.json", tok) // Сохраните полученный токен в файл
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func createFolder(service *drive.Service, folderName string) {
	folder := &drive.File{
		Name:     folderName,
		MimeType: "application/vnd.google-apps.folder",
	}
	_, err := service.Files.Create(folder).Do()
	if err != nil {
		log.Fatalf("Не удалось создать папку: %v", err)
	}
	log.Println("Папка создана успешно!")
}

func main() {
	// Чтение содержимого JSON-файла
	data, err := os.ReadFile("D:\\programming\\go\\netMg\\netMg\\token.json")
	if err != nil {
		log.Fatalf("Не удалось прочитать JSON-файл: %v", err)
	}
	fmt.Println(data)
	credentials := string(data)
	config, err := google.ConfigFromJSON([]byte(credentials), drive.DriveFileScope)
	if err != nil {
		log.Fatalf("Не удалось загрузить данные клиента: %v", err)
	}

	client := getClient(config)

	service, err := drive.NewService(client)
	if err != nil {
		log.Fatalf("Не удалось создать клиент Google Drive: %v", err)
	}

	createFolder(service, "Название папки")
}
