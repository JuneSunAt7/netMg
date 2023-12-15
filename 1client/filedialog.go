package client

import (
	"log"

	"github.com/harry1453/go-common-file-dialog/cfd"
)

func ChooseFile() string {
	openDialog, err := cfd.NewOpenFileDialog(cfd.DialogConfig{
		Title:                   "Open A File",
		Role:                    "OpenFile",
		SelectedFileFilterIndex: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := openDialog.Show(); err != nil {
		log.Fatal(err)
	}
	result, err := openDialog.GetResult()
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	return result
}
