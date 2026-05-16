package main

import (
	"effective_mobile/config"
	"effective_mobile/internal/sub"
	"effective_mobile/internal/migrate"
	"effective_mobile/logger"
	"effective_mobile/pkg/db"
	"effective_mobile/pkg/middleware"
	"net/http"

	"github.com/rs/zerolog/log"
)

func App(conf *config.DatabaseConfig) http.Handler {

	db := db.NewDB(conf)
	migrate.Migrate(db.DB)

	router := http.NewServeMux()

	//Repositories
	subRepository := sub.NewSubscriptionRepository(db)
	
	//Services
	subService := sub.NewSubscriptionService(subRepository)

	// Handler
	sub.NewSubscriptionHandler(router, &sub.SubscriptionHandlerDeps{
		SubRepo: subRepository,
		SubService: subService,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}

func main() {
	logger.Init(config.DefaultLogConfig())
	
	// Config
	if err := config.Init(); err != nil {
		log.Warn().Err(err).Msg("failed to load .env, using environment variables")
	} else {
		log.Info().Msg(".env file loaded")
	}	
	// Logger
	logger.Init(config.NewLogConfig())

	app := App(config.NewDatabaseConfig())
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	log.Info().
		Msg("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Server failed to start")
	}
}
