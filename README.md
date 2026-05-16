# Effective_mobile тестовое

# Структура проекта
```md
├── 📁 cmd
    ├── 📁 migrations - миграции для БД
        ├── 📄auto.go - ручной запуск миграций
    ├── 📄main.go - основное приложение
├── 📁 config
    ├── 📄config.go - функционал для загрузки системных переменных из .env файла
├── 📁 docker - всё связанное с докером   
├── 📁 internal - внутренний код для реализации работы подписками, валидацией запросов и авто-миграции при запуске приложения
    ├── 📁 sub  - модуль работы с подписками
    ├── 📁 migrate
    ├── 📁 validate
├── 📁 logger - реализация глобального логгера
├── 📁 pkg - Код переиспользуемых библиотек(работа с БД, middleware, request, response)
├── 📄 LICENSE
```
# Файл .env
Пример для .env файла. Файл размещается в корне проекта.

```.env
DB_HOST="postgres"
DB_PORT="5432"
POSTGRES_USER="mega_user"
POSTGRES_PASSWORD="1111"
POSTGRES_DB="subscriptions"
DB_SSLMODE="disable"

LOG_LEVEl=0
LOG_LEVEL_STR="debug"
LOG_FORMAT="json"
```

ВАЖНО!
DB_HOST="postgres" обязательно, если проект запускаем в докере.
Если запускаем на локальной тачке, то DB_HOST="localhost".

# Запуск и сборка проекта

Всё сделано через mage

## Установка mage
```cmd
go install github.com/magefile/mage@latest
go get github.com/magefile/mage/sh
```
### Команды

| Команда | Описание |
|---|---|
| `mage up` | Запустить все сервисы в фоне |
| `mage down` | Остановить все сервисы |
| `mage build` | Пересобрать и запустить все сервисы |
| `mage migrate` | Запустить вручную миграции базы данных

Миграции в проекте происходят автоматически при запуске проекта. вручную их запускать не обязательно.

### Быстрый старт

```bash
# 1. Скопируй .env.example в .env и заполни переменные
cp .env.example .env

# 2. Собрать и запустить все сервисы в docker
mage build
```

### Сервисы

| Сервис | Адрес |
|---|---|
| API | http://localhost:8081 |
| Swagger UI | http://localhost:8082 |
| PostgreSQL | localhost:5432 |



