package api

import (
	"fmt"
	"struct/config"
)

func SomeApiFunc() {
	fmt.Println("Api func")
	cfg := config.NewConfig()
	fmt.Println(cfg)
}
