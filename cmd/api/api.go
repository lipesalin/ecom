package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lipesalin/service/user"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func newAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (server *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	useHandler := user.NewHandler()
	useHandler.RegisterRoutes()

	log.Println("Servidor online", server.address)

	return http.ListenAndServe(server.address, router)
}
