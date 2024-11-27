package api

import (
	"log"
	"net/http"
	"rest-api-demo/routes"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Address string
}

func NewAPIServer(address string) *APIServer{
	return &APIServer{
		Address: address,
	}
}

func (s *APIServer) Run() error {
	// create router
	router := mux.NewRouter()

	// register services
	handler := routes.NewHandlerNoParam()
	handler.RegisterRoutes(router)
	log.Println("Listening on port", s.Address)

	err := http.ListenAndServe(s.Address, router)
	if (err != nil){
		log.Fatalf("error when run server: %s", err)
	}
	return nil
}