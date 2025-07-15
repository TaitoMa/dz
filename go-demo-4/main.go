package main

import (
	"fmt"
	"go-demo-4/account"
)

func main() {
	fmt.Println("___Менеджер паролей___\n")
	vault := account.NewVault()

Menu:
	for {
		menuItem := getMenu()
		switch menuItem {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func deleteAccount(vault *account.Vault) {

}

func findAccount(vault *account.Vault) {
	fmt.Println("Введите URL для поиска")
	var url string
	fmt.Scan(&url)

	accounts := vault.FindAccount(url)

	for _, acc := range accounts {
		acc.OutputPassword()
	}
}

func getMenu() int {
	fmt.Println("Выберите вариант")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil || n < 1 || n > 4 {
			fmt.Println("Ошибка ввода")
			continue
		}
		break
	}

	return n
}

func createAccount(vault *account.Vault) {
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

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
