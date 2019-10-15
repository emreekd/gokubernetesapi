package api

import (
	"net/http"
	"strings"

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
		r.Get("/*", h.defaultHandler)
	})

	return r
}

func (h *portcontrollerhandler) defaultHandler(w http.ResponseWriter, r *http.Request) {

	var pathInfo = strings.Split(r.RequestURI, "/")
	if pathInfo != nil && len(pathInfo) > 2 {
		resp := h.kubeService.GetByNamespace(pathInfo[len(pathInfo)-1])
		render.JSON(w, r, resp)
	}
}
