package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	r *chi.Mux // r inhibits all the methods presten on type *chi.Mux(router)
}

func NewServer() *Server {
	router := chi.NewRouter()
	return &Server{r: router}
}

// Start starts the server on the port
// passed in
func (s *Server) Start(port string) {
	if err := http.ListenAndServe(":"+port, s.r); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
