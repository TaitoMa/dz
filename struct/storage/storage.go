package storage

import (
	"github.com/fatih/color"
	"struct/files"
)

func SaveFile(content []byte, name string) {
	err := files.WriteFile(content, name)
	if err != nil {
		color.Red(err.Error(), "SaveFile")
	}
}

func ReadFile(name string) ([]byte, error) {
	content, err := files.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return content, nil
}
