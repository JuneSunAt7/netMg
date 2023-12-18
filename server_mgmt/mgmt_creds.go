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
	options = append(options, fmt.Sprintf("Выход"))
	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.TextStyle.Add(*pterm.NewStyle(pterm.FgBlue))
	printer.KeyConfirm = keys.Enter

	selectedOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if selectedOptions == "Выход" {
		return
	} else {
		pterm.Info.Printfln("Выбранный пользователь: %s", pterm.Green(selectedOptions))
		delUserinCreds(selectedOptions)
	}

}
func delUserinCreds(key string) {
	file, err := os.Open("user_creds.db")
	if err != nil {
		pterm.Error.Println("Ошибка при открытии файла")
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		pterm.Error.Println("Ошибка при чтении файла")
		return
	}

	var records []MyStruct
	err = json.Unmarshal(data, &records)
	if err != nil {
		pterm.Error.Println("Ошибка при парсинге")
		return
	}

	records = removeRecord(records, key)

	newData, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		pterm.Error.Println("Ошибка при кодировании данных")
		return
	}

	err = ioutil.WriteFile("user_creds.db", newData, 0644)
	if err != nil {
		pterm.Error.Println("Ошибка при записи в файл")
		return
	}
	pterm.Success.Println("Пользователь ", key, " успешно удален")
}
func removeRecord(records []MyStruct, name string) []MyStruct {
	var updatedRecords []MyStruct
	for _, record := range records {
		if record.UserName != name {
			updatedRecords = append(updatedRecords, record)
		}
	}
	return updatedRecords
}
func UserData() {

}
func ConfigUser() {

}
