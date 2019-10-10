package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	router chi.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func New() *Server {
	s := &Server{}

	r := chi.NewRouter()
	r.Route("/port", func(r chi.Router) {
		h := portcontrollerhandler{}
		r.Mount("/", h.router())
	})

	s.router = r

	return s
}
