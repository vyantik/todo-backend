# Todo Backend

Бэкенд-сервис для управления задачами (Todo), написанный на Go.

## Технологии

-   Go 1.24.2
-   Gin Web Framework
-   PostgreSQL
-   JWT для аутентификации
-   Docker & Docker Compose
-   Swagger для API документации

## Требования

-   Go 1.24.2 или выше
-   Docker и Docker Compose
-   PostgreSQL (если запускаете локально)

## Установка и запуск

1. Клонируйте репозиторий:

```bash
git clone https://github.com/vyantik/todo-backend.git
cd todo-backend
```

2. Создайте файл .env в корневой директории:

```bash
cp .env.example .env
```

3. Запустите с помощью Docker Compose:

```bash
docker-compose up -d
```

Или запустите локально:

```bash
make run
```

## API Документация

После запуска сервера, Swagger документация доступна по адресу:

```
http://localhost:8080/swagger/index.html
```

## Структура проекта

```
.
├── cmd/            # Точка входа приложения
├── configs/        # Конфигурационные файлы
├── docs/           # Документация
├── pkg/            # Внутренние пакеты
├── schema/         # SQL схемы
├── docker-compose.yaml
├── go.mod
├── go.sum
├── Makefile
└── server.go
```

## Команды Makefile

-   `make run` - Запуск приложения
-   `make build` - Сборка приложения
-   `make test` - Запуск тестов

## Лицензия

MIT
