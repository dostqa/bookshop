package main

import (
	"bookshop/internal/config"
	"bookshop/internal/handlers/books"
	"bookshop/internal/logger"
	"bookshop/internal/logger/utils"
	"bookshop/internal/middleware"
	"bookshop/internal/storage"
	"log"
	"net/http"
	"os"

	mw "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	config, err := config.NewConfigFromFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("Чтение конфигурационного файла не было выполнено:\n %v", err)
	}

	logger := logger.NewLogger(config.Env)
	logger.Info("Bookshop запущен")

	storage, err := storage.NewStorage(config.StoragePath)
	if err != nil {
		logger.Error("Не удалось подключиться к базе данных:", utils.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	// Если нужно логировать id запросов:
	// router.Use(middleware.RequestID)
	router.Use(middleware.NewMiddlewareLogger(logger))
	router.Use(mw.Recoverer)

	router.Get("/api/books", books.GetAll(logger, storage))

	srv := &http.Server{
		Addr:         config.Address,
		Handler:      router,
		ReadTimeout:  config.HTTPServer.Timeout,
		WriteTimeout: config.HTTPServer.Timeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Не удалось запустить сервер", utils.Err(err))
	}

	logger.Error("Сервер остановил свою работу")

	// 4. Написать хендлеры, для каждого запроса
	// router.Get("/books", booksHandler.Get)
	// router.Post("/books", booksHandler.Create)
	// 4.0 Получение всех книг GET /api/books books.GetAll

	// 5. Написать роутер
	// 6. Написать сервер и запустить его.
}

// curl --request GET \
//   --url http://localhost:8081/api/books \
//   --header 'User-Agent: insomnia/12.5.0'
