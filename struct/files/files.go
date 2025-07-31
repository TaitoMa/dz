package files

import (
	"errors"
	"github.com/fatih/color"
	"os"
	"strings"
)

type JsonDb struct {
	Filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		Filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	if !strings.HasSuffix(db.Filename, ".json") {
		color.Red("Файл должен быть .json")
		return nil, errors.New("file name not contains .json")
	}
	data, err := os.ReadFile(db.Filename)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) error {
	file, err := os.Create(db.Filename)
	if err != nil {
		color.Red(err.Error(), "Writing"+db.Filename)
		return err
	}
	_, err = file.Write(content)
	if err != nil {
		color.Red(err.Error(), "Writing"+db.Filename)
		return err
	}
	color.Green("Файл записан")
	defer file.Close()
	return nil
}
