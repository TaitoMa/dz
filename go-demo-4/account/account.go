package account

import (
	"errors"
	"github.com/fatih/color"
	"math/rand/v2"
	"time"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	//_, err := url.ParseRequestURI(urlString)
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	//if err != nil {
	//	return nil, errors.New("INVALID_URL")
	//}
	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func (acc *Account) OutputPassword() {
	color.Cyan("Password: " + acc.Password)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY1234")
