package main

import (
	"log"

	"github.com/duixe/social_app/internal/db"
	"github.com/duixe/social_app/internal/env"
	"github.com/duixe/social_app/internal/repository"
)

func main() {

	addr := env.Envs.DBAddress
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	
	repository := repository.NewRepository(conn)
	db.Seed(repository, conn)
}