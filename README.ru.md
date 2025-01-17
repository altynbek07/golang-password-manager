# Менеджер Паролей

Простой менеджер паролей, написанный на Go.

## Возможности

- Создание, поиск и удаление аккаунтов
- Шифрование и дешифрование данных аккаунта
- Хранение данных локально в JSON файле

## Требования
- Go (версия указана в `go.mod`)

## Установка

1. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/altynbek07/golang-password-manager.git
    ```

2. Перейдите в директорию проекта:
   ```sh
   cd golang-password-manager
   ```

3. Установите зависимости:
    ```sh
    go mod tidy
    ```

4. Создайте файл `.env` на основе `.env.example`:
    ```sh
    cp .env.example .env
    ```

5. Установите ваш секретный ключ в файле `.env` (вы можете сгенерировать секретный **Encryption key 256** и скопировать [здесь](https://acte.ltd/utils/randomkeygen)):
    ```env
    KEY=your_secret_key
    ```

6. Соберите проект:
   ```sh
   go build -o password-manager
   ```

7. Запустите приложение:
   ```sh
   ./password-manager
   ```

## Использование

После запуска приложения вы можете взаимодействовать с ним через консоль:

- **Создать аккаунт**:
  Введите имя, пароль и URL при запросе.
- **Найти аккаунт по URL**:
  Приложение покажет все аккаунты, содержащие переданный URL.
- **Найти аккаунт по логину**:
  Приложение покажет все аккаунты, содержащие переданный логин.
- **Удалить аккаунт**:
  Введите URL аккаунта, который хотите удалить.

## Пример

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

## Автор
Altynbek Kazezov

## Лицензия
Этот проект лицензирован по лицензии MIT. См. файл [LICENSE](LICENSE) для получения подробной информации.
