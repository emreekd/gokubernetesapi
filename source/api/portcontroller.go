package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type portcontrollerhandler struct {
}

func (h *portcontrollerhandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/service", func(r chi.Router) {
		r.Get("/", h.defaultHandler)
	})

	return r
}

func (h *portcontrollerhandler) defaultHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "this is port service")
}
