package storage

import (
	"struct/files"
)

func SaveFile(content []byte, name string) {
	files.WriteFile(content, name)
}

func ReadFile(name string) []byte {
	content, _ := files.ReadFile(name)

	return content
}
