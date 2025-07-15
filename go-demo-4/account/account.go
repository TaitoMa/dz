package account

import (
	"errors"
	"github.com/fatih/color"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &AccountWithTimeStamp{
		Account: Account{
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

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func (acc *Account) OutputPassword() {
	color.Cyan("Password: " + acc.password)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY1234")
