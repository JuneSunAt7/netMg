package servermgmt

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"

	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
)

type MyStruct struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
type CredArr []MyStruct

func (p *CredArr) FromJSON(r io.Reader) error {
	en := json.NewDecoder(r)
	return en.Decode(p)
}
func GetCred() (*CredArr, error) {
	f, err := os.Open("user_creds.db")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var creds CredArr
	err = creds.FromJSON(f)
	if err != nil {
		return nil, err
	}

	return &creds, nil
}
func AddUser() {
	uname, _ := pterm.DefaultInteractiveTextInput.Show("Имя")
	passwd, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Пароль")

	filename := "user_creds.db"
	file, err := os.ReadFile(filename)
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

	err = os.WriteFile(filename, dataBytes, fs.FileMode(0644))
	if err != nil {
		fmt.Println(err)
	}

}
func DelUser() {
	creds, err := GetCred()
	if err != nil {
		pterm.Error.Println(err)
		return
	}
	var options []string
	for _, cred := range *creds {
		options = append(options, fmt.Sprintf(cred.UserName))

	}
	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.TextStyle.Add(*pterm.NewStyle(pterm.FgBlue))
	printer.KeyConfirm = keys.Enter

	selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	pterm.Info.Printfln("Выбранный пользователь: %s", pterm.Green(selectedOptions))
	delUserinCreds(selectedOptions)

}
func delUserinCreds(key string) {
	// TODO #16 fix function: parse db file
	file, err := os.OpenFile("user_creds.db", os.O_RDWR, 0644)
	if err != nil {
		pterm.Error.Println("Невозможно открыть файл")
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		pterm.Error.Println("Невозможно прочитать файл")

	}

	// Структура для разбора данных JSON
	yourData := map[string]interface{}{"username": "John", "password": 30}
	err = json.Unmarshal(data, &yourData)
	if err != nil {
		pterm.Error.Println("Невозможно спарсить файл")
	}

	// Удаление записи из переменной yourData
	delete(yourData, key)

	// Запись обновленных данных в файл
	file, err = os.OpenFile("user_creds.db", os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		pterm.Error.Println("Невозможно открыть файл для записи")
	}
	defer file.Close()

	// Преобразование обновленных данных обратно в JSON и запись в файл
	newData, err := json.Marshal(yourData)
	if err != nil {
		pterm.Error.Println("Невозможно открыть файл")
	}

	_, err = file.Write(newData)
	if err != nil {
		pterm.Error.Println("Невозможно записать в файл")
	}
}
func UserData() {

}
func ConfigUser() {

}
