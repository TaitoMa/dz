package api

import (
	"fmt"
	"struct/config"
)

type Methods struct {
}

func SomeApiFunc() {
	fmt.Println("Api func")
	cfg := config.NewConfig()
	fmt.Println(cfg)
}
