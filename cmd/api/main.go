package main

import (
	"log"

	"github.com/duixe/social_app/internal/db"
	"github.com/duixe/social_app/internal/env"
	"github.com/duixe/social_app/internal/repository"
)

const version = "1.1.0"

//	@title			social APP API
//	@version		1.0
//	@description	This is simple social media developed whiles following a udemy course.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/v1
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description
func main() {

	cfg := config{
		addr: env.Envs.Port,
		apiURL: env.Envs.APPUrl,
		db: dbConfig{
			addr:         env.Envs.DBAddress,
			maxOpenConns: int(env.GetInt("DB_MAX_OPEN_CONNS", 30)),
			maxIdleConns: int(env.GetInt("DB_MAX_IDLE_CONNS", 30)),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
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
