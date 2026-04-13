package books

import (
	"bookshop/internal/logger/utils"
	"bookshop/internal/models"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	// "github.com/go-chi/chi/middleware"
)

func ToBook(book models.Book) Book {
	return Book{
		Title:    book.Title,
		Category: book.Category,
		Author:   book.Author,
	}
}

type Book struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Author   string `json:"author"`
}

// TODO: Выделить для данных структур отдельный пакет

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	Books []Book `json:"books"`
	Error *Error `json:"error,omitempty"`
}

func newOK(books []Book) Response {
	return Response{
		Books: books,
	}
}

func newError(message string) Response {
	return Response{
		Error: &Error{
			Message: message,
		},
	}
}

type BookGetter interface {
	GetAllBooks() ([]models.Book, error)
}

func GetAll(logger *slog.Logger, bookGetter BookGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.books.GetAll"

		log := logger.With(
			slog.String("op", op),
			// Если нужно логировать id запросов:
			// slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		res, err := bookGetter.GetAllBooks()
		if err != nil {
			log.Error("Не удалось получить книги по запросу:", utils.Err(err))

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, newError("Internal server error"))

			return
		}

		books := make([]Book, 0, len(res))
		for _, book := range res {
			books = append(books, ToBook(book))
		}

		log.Info("Получены книги по запросу")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, newOK(books))
	}
}
