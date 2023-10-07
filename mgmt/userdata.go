package mgmt

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func compress(source string, target string) error {
	outFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Относительный путь внутри архива
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		// Создание нового файла или папки в архиве
		if info.IsDir() {
			zipWriter.CreateHeader(&zip.FileHeader{
				Name:     relPath + "/",
				Method:   zip.Store, // без сжатия
				Modified: info.ModTime(),
			})
		} else {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Создание нового файла в архиве
			zipFile, err := zipWriter.CreateHeader(&zip.FileHeader{
				Name:     relPath,
				Method:   zip.Deflate, // сжатие
				Modified: info.ModTime(),
			})
			if err != nil {
				return err
			}

			// Запись содержимого файла в архив
			_, err = io.Copy(zipFile, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
func CreateContainer(path string) {
	contName := time.Now().Format(time.DateTime) + ".zip"
	err := compress(path, contName)
	if err != nil {
		fmt.Println("Ошибка сжатия:", err)
	} else {
		fmt.Println("Сжатие выполнено успешно!")
	}
}
