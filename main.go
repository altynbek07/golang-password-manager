package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		choice := getMenu()
		switch choice {
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

func getMenu() int {
	var choice int

	fmt.Println("Выберите меню:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&choice)

	return choice
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль (можете оставить пустым)")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}

	vault.AddAccount(*myAccount)

}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrl(url)

	if len(accounts) == 0 {
		color.Yellow("Аккаунты не найдены")
		return
	}

	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Yellow("Не найдено")
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}
