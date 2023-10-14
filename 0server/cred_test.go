package server

import (
	"os"
	"testing"
)

func TestCredentials_FromJSON(t *testing.T) {
	f,_ := os.Open("test.json")
	var creds CredArr
	err := creds.FromJSON(f)
	if err !=nil{
		t.Fatal(err)
	}

	for _,obj := range creds{
		t.Log(obj.Password, " ", obj.Username) 
	}
}
