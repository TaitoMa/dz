package main

import (
	"demo/app-1/catsApp/cat"
	"demo/app-1/catsApp/translate"
	"fmt"
)

func main() {
	for {
		fmt.Println("Вот вам факт: ")
		ct := cat.GetCatFact()
		fmt.Println(ct.Fact)
		res, err := translate.GetTranslate(ct.Fact)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res)
		fmt.Println("Если хотите выйти введите 0: ")
		var s string
		fmt.Scan(&s)
		if s == "0" {
			break
		}
	}
}
