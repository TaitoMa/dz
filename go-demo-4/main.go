package main

import (
	"fmt"
	"go-demo-4/account"
)

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	promptUrl := promptData("Введите URL")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, promptUrl)
	if err != nil {
		fmt.Println("Неверный формат урла или логина")
		return
	}
	myAccount.OutputPassword()
	fmt.Println(*myAccount)
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
