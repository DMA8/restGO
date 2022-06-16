package entrypoint

import (
	"testTask/internal/usecases"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type useCaser interface {
	
}

type Handler struct {
	useCase *usecases.UseCase
}

func NewHandler(usecase *usecases.UseCase) *Handler {
	return &Handler{
		useCase: usecase,
	}
}

func NewRouter(h *Handler) *chi.Mux {
	newRouter := chi.NewRouter()
	newRouter.Use(middleware.Logger)
	newRouter.Get("/id", h.getByIDs)
	newRouter.Post("/create", h.create)
	newRouter.Post("/create1", h.create1)
	newRouter.Post("/update", h.updateProps)
	return newRouter
}
