package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lipesalin/ecom/cmd/api"
	"github.com/lipesalin/ecom/config"
	"github.com/lipesalin/ecom/db"
)

func main() {
	// config da conexão
	psqlconn :=
		fmt.Sprintf(`
			host=%s 
			port=%s 
			user=%s 
			password=%s 
			dbname=%s 
			sslmode=disable`,
			config.Envs.DBHost,
			config.Envs.DBPort,
			config.Envs.DBUser,
			config.Envs.DBPassword,
			"ecom",
		)

	// conexão
	db, err := db.NewDBConnection(psqlconn)

	initStorage(db)

	if err != nil {
		log.Fatal("Erro ao tentar conectar no banco.", err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Banco: conectado.")
}
