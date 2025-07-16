package files

import (
	"errors"
	"github.com/fatih/color"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if !strings.Contains(name, ".json") {
		color.Red("Файл должен быть .json")
		return nil, errors.New("file name not contains .json")
	}
	data, err := os.ReadFile(name)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		color.Red(err.Error(), "Writing"+name)
		return
	}
	_, err = file.Write(content)
	if err != nil {
		color.Red(err.Error(), "Writing"+name)
		return
	}
	color.Green("Файл записан")
	defer file.Close()
}
