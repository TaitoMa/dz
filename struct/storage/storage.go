package storage

import (
	"errors"
	"github.com/fatih/color"
	"os"
	"strings"
)

type StorageDb struct {
	Prefix string
}

func NewStorage(prefix string) *StorageDb {
	return &StorageDb{
		Prefix: prefix,
	}
}

func (db *StorageDb) Read() ([]byte, error) {
	if !strings.HasSuffix(db.Prefix, ".json") {
		color.Red("Файл должен быть .json")
		return nil, errors.New("file name not contains .json")
	}
	data, err := os.ReadFile(db.Prefix)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	return data, nil
}

func (db *StorageDb) Write(content []byte) error {
	file, err := os.Create(db.Prefix)
	if err != nil {
		color.Red(err.Error(), "Writing"+db.Prefix)
		return err
	}
	_, err = file.Write(content)
	if err != nil {
		color.Red(err.Error(), "Writing"+db.Prefix)
		return err
	}
	color.Green("Файл записан")
	defer file.Close()
	return nil
}
