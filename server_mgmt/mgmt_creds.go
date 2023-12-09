package servermgmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pterm/pterm"
)

type MyStruct struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func AddUser() {
	uname, _ := pterm.DefaultInteractiveTextInput.Show("Имя")
	passwd, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Пароль")

	filename := "user_creds.db"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	data := []MyStruct{}

	// Here the magic happens!
	json.Unmarshal(file, &data)

	newStruct := &MyStruct{
		UserName: uname,
		PassWord: passwd,
	}

	data = append(data, *newStruct)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

}
func DelUser() {

}
func UserData() {

}
func ConfigUser() {

}
