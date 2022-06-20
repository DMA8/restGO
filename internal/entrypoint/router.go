package entrypoint

import (
	"testTask/internal/domain"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type useCaser interface {
	GetByIDs([]uint) ([]*domain.Prop, error)
	CreateProps(*domain.Props) error
	UpdateProps(*domain.Props) error
}

type Handler struct {
	useCase useCaser
}

func NewHandler(usecase useCaser) *Handler {
	return &Handler{
		useCase: usecase,
	}
}

func NewRouter(h *Handler) *chi.Mux {
	newRouter := chi.NewRouter()
	newRouter.Use(middleware.Logger)
	newRouter.Use(h.validateQuery)
	newRouter.Get("/id", h.getByIDs)
	newRouter.Post("/create", h.create)
	newRouter.Post("/update", h.updateProps)
	return newRouter
}
