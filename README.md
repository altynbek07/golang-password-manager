# Password Manager

A simple password manager written in Go.

## Features

- Create, find, and delete accounts
- Encrypt and decrypt account data
- Store data locally in a JSON file

## Requirements
- Go (version specified in `go.mod`)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/altynbek07/golang-password-manager.git
    ```

2. Navigate to the project directory:
   ```sh
   cd golang-password-manager
   ```

3. Install dependencies:
    ```sh
    go mod tidy
    ```

4. Create a `.env` file based on the `.env.example`:
    ```sh
    cp .env.example .env
    ```

5. Set your secret key in the `.env` file:
    ```env
    KEY=your_secret_key
    ```

6. Build the project:
   ```sh
   go build -o password-manager
   ```

7. Run the application:
   ```sh
   ./password-manager
   ```

## Usage

After running the application, you can interact with it through the console:

- **Create account**:
  Enter the name, password and URL when prompted.
- **Find account by URL**:
  The application will show all accounts that contain the passed URL
- **Find account by login**:
  The application will show all accounts that contain the passed login
- **Delete account**:
  Enter the URL of the account you want to delete.

## Example

```bash
__Менеджер паролей__
Найдено 0 аккаунтов
1. Создать аккаунт
2. Найти аккаунт по URL
3. Найти аккаунт по логину
4. Удалить аккаунт
5. Выход
Выберите меню: 1
Введите логин: user_login
Введите пароль (можете оставить пустым): qwerty
Введите URL: https://google.com

Выберите меню: 2
Введите URL для поиска: google.com
user_login
qwerty
https://google.com

Выберите меню: 3
Введите логин для поиска: user_login
user_login
qwerty
https://google.com

Выберите меню: 4
Введите URL для удаления: https://google.com
```

## Author
Altynbek Kazezov

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.