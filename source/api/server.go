package api

import (
	"net/http"

	"../services"
	"github.com/go-chi/chi"
)

type Server struct {
	router chi.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s.router.ServeHTTP(w, r)
}

func New(kubeService services.IKubeService) *Server {
	s := &Server{}

	r := chi.NewRouter()
	r.Route("/kube", func(r chi.Router) {
		h := kubecontrollerhandler{
			kubeService: kubeService,
		}
		r.Mount("/", h.router())
	})
	r.Route("/pod", func(r chi.Router) {
		h := podcontrollerhandler{
			kubeService: kubeService,
		}
		r.Mount("/", h.router())
	})

	s.router = r

	return s
}
