package main

import (
	"log"

	"github.com/duixe/social_app/internal/db"
	"github.com/duixe/social_app/internal/env"
	"github.com/duixe/social_app/internal/repository"
)

func main() {

	cfg := config{
		addr: env.Envs.Port,
		db: dbConfig{
			addr: env.Envs.DBAddress,
			maxOpenConns: int(env.GetInt("DB_MAX_OPEN_CONNS", 30)),
			maxIdleConns: int(env.GetInt("DB_MAX_IDLE_CONNS", 30)),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),	
		},
	}

	//instantiate db conn
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection is successful")

	//instatiate the repository by passing in the instantiated db
	repository := repository.NewRepository(db)

	app := &application{
		config:     cfg,
		repository: repository,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
