package main

import (
	"fmt"
	"go-demo-4/account"
	"go-demo-4/cloud"
	"go-demo-4/files"
)

func main() {
	fmt.Println("___Менеджер паролей___\n")
	db := files.NewJsonDb("data.json")
	cloudDb := cloud.NewCloudDb("https://check.com")
	vault := account.NewVault(db)
	vault1 := account.NewVault(cloudDb)
	fmt.Println(vault1)

Menu:
	for {
		menuItem := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		switch menuItem {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func deleteAccount(vault *account.VaultWithDb) {

}

func findAccount(vault *account.VaultWithDb) {
	fmt.Println("Введите URL для поиска")
	var url string
	fmt.Scan(&url)

	accounts := vault.FindAccount(url)

	for _, acc := range accounts {
		acc.OutputPassword()
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	promptUrl := promptData([]string{"Введите URL"})
	myAccount, err := account.NewAccount(login, password, promptUrl)
	if err != nil {
		fmt.Println("Неверный формат урла или логина")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData[T any](prompt []T) string {
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
