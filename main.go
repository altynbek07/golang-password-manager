package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите ссылку: ")

	myAccount, err := account.NewAccountWithTimestamp(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}

	myAccount.OutputPassword()

	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}
