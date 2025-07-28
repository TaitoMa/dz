package main

import (
	"fmt"
	"go-demo-4/account"
	"go-demo-4/cloud"
	"go-demo-4/files"
	"strings"
)

var menu = map[string]func(db *account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по Login",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func menuCounter() func() {
	count := 0
	return func() {
		count++
		fmt.Println(count)
	}
}

func main() {
	fmt.Println("___Менеджер паролей___\n")
	db := files.NewJsonDb("data.json")
	cloudDb := cloud.NewCloudDb("https://check.com")
	vault := account.NewVault(db)
	vault1 := account.NewVault(cloudDb)
	fmt.Println(vault1)
	counter := menuCounter()
Menu:
	for {
		counter()
		menuItem := promptData(menuVariants...)
		menuFunc := menu[menuItem]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func deleteAccount(vault *account.VaultWithDb) {

}

func findAccount(vault *account.VaultWithDb) {
	fmt.Println("Введите URL для поиска")
	var url string
	fmt.Scan(&url)

	accounts := vault.FindAccount(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	for _, acc := range accounts {
		acc.OutputPassword()
	}
}
func findAccountByLogin(vault *account.VaultWithDb) {
	fmt.Println("Введите Login для поиска")
	var login string
	fmt.Scan(&login)

	accounts := vault.FindAccount(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	for _, acc := range accounts {
		acc.OutputPassword()
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	promptUrl := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, promptUrl)
	if err != nil {
		fmt.Println("Неверный формат урла или логина")
		return
	}
	vault.AddAccount(*myAccount)
}

// func promptData[T any](prompt []T) string {
func promptData(prompt ...string) string {
	for i, p := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", p)
		} else {
			fmt.Println(p)
		}
	}
	var res string
	fmt.Scan(&res)
	return res
}
