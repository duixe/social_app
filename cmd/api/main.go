package main

import (
	"log"

	"github.com/duixe/social_app/internal/env"
)

func main()  {

	cfg := config {
		addr: env.Envs.Port,
	}
	
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}