package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите меню",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Ошибка при загрузке .env файла")
	}

	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
Menu:
	for {
		choice := promptData(menuVariants...)

		menuFunc := menu[choice]

		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}

}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль (можете оставить пустым)")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или логин")
		return
	}

	vault.AddAccount(*myAccount)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(account account.Account, query string) bool {
		return strings.Contains(account.Url, query)
	})

	outputFindAccountResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(account account.Account, query string) bool {
		return strings.Contains(account.Login, query)
	})

	outputFindAccountResult(&accounts)
}

func outputFindAccountResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Yellow("Аккаунты не найдены")
		return
	}

	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func promptData(prompt ...string) string {
	for index, value := range prompt {
		if index == len(prompt)-1 {
			fmt.Printf("%v: ", value)
		} else {
			fmt.Println(value)
		}
	}
	var result string
	fmt.Scanln(&result)
	return result
}
