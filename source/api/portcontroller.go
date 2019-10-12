package api

import (
	"net/http"

	"../services"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type portcontrollerhandler struct {
	kubeService services.IKubeService
}

func (h *portcontrollerhandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/service", func(r chi.Router) {
		r.Get("/", h.defaultHandler)
	})

	return r
}

func (h *portcontrollerhandler) defaultHandler(w http.ResponseWriter, r *http.Request) {
	resp := h.kubeService.GetAllPods()
	render.JSON(w, r, resp)
}
