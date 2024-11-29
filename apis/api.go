package api

import (
	"database/sql"
	"log"
	"net/http"
	"rest-api-demo/routes"
	"rest-api-demo/services/users"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Address string
	DbConnection *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer{
	return &APIServer{
		Address: address,
		DbConnection: db,
	}
}

func (s *APIServer) Run() error {
	// create router
	router := mux.NewRouter()

	// register services
	userStore := services.NewUserStore(s.DbConnection)

	handler := routes.NewHandler(userStore)
	handler.RegisterRoutes(router)
	log.Println("Listening on port", s.Address)

	err := http.ListenAndServe(s.Address, router)
	if (err != nil){
		log.Fatalf("error when run server: %s", err)
	}
	return nil
}