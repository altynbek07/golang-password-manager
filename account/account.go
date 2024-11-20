package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimestamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.password)
	// fmt.Println(acc.password)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for index := range res {
		res[index] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimestamp(login, password, urlString string) (*AccountWithTimestamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimestamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login:    login,
			url:      urlString,
			password: password,
		},
	}

	if password == "" {
		newAcc.generatePassword(16)
	}

	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")
