package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	createAccount()
}

func createAccount() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите ссылку: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}

	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать данные в JSON")
		return
	}

	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}
