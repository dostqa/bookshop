# Bookshop

Bookshop — это небольшой учебный проект для изучения основ REST API.

## Текущее состояние проекта

На данный момент реализован один базовый эндпоинт для получения всех книг из базы данных: **GET /api/books**

### Пример запроса

```bash
curl --request GET \
  --url http://localhost:8081/api/books
```

## Планируемый функционал

### Рядом указаны примеры планируемых эндпоинтов
* Получение книги по названию
  `GET /api/books?title=harry+potter`

* Получение книги по id
  `GET /api/books/{id}`

* Добавление книги
  `POST /api/books`

* Обновление книги
  `PUT /api/books/{id}`

* Удаление книги
  `DELETE /api/books/{id}`

## Используемые технологии и зависимости

### Go

* net/http
* log/slog

### Роутер

* github.com/go-chi/chi/v5

### База данных

* SQLite (github.com/mattn/go-sqlite3)