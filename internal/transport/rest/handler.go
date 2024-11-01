package rest

import (
	"context"
	"net/http"

	"github.com/asliddinberdiev/crud_basic/internal/domain"
	"github.com/gorilla/mux"
)

type Books interface {
	Create(ctx context.Context, book domain.Book) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, input domain.UpdateBookInput) error
}

type Handler struct {
	booksService Books
}

func NewHandler(books Books) *Handler {
	return &Handler{booksService: books}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	books := r.PathPrefix("/books").Subrouter()
	{
		books.HandleFunc("", h.createBook).Methods(http.MethodPost)
		books.HandleFunc("", h.getAllBooks).Methods(http.MethodGet)
		books.HandleFunc("/{id: [0-9]+}", h.getBookByID).Methods(http.MethodGet)
		books.HandleFunc("/{id: [0-9]+}", h.deleteBook).Methods(http.MethodDelete)
		books.HandleFunc("/{id: [0-9]+}", h.updateBook).Methods(http.MethodPut)
	}

	return r
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request)  {}
