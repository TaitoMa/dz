package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

//func newAccount(login, password, urlString string) (*account, error) {
//	_, err := url.ParseRequestURI(urlString)
//	if login == "" {
//		return nil, errors.New("EMPTY_LOGIN")
//	}
//	if err != nil {
//		return nil, errors.New("INVALID_URL")
//	}
//	newAcc := &account{login: login, password: password, url: urlString}
//	if password == "" {
//		newAcc.generatePassword(12)
//	}
//
//	return newAcc, nil
//}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStamp{
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func (acc *account) outputPassword() string {
	return acc.password
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY1234")

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	promptUrl := promptData("Введите URL")
	myAccount, err := newAccountWithTimeStamp(login, password, promptUrl)
	if err != nil {
		fmt.Println("Неверный формат урла или логина")
		return
	}
	fmt.Println(*myAccount)
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
