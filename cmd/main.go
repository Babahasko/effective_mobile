package main

import (
	"effective_mobile/config"
	"effective_mobile/internal/sub"
	"effective_mobile/logger"
	"effective_mobile/pkg/db"
	"net/http"
)

func App(conf *config.DatabaseConfig) http.Handler {

	db := db.NewDB(conf)
	router := http.NewServeMux()

	//Repositories
	subRepository := sub.NewSubscriptionRepository(db)
	// statRepository := stat.NewStatRepository(db) // пока под вопросом
	//Services
	subService := sub.NewSubscriptionService(subRepository)

	// Handler
	sub.NewSubscriptionHandler(router, &sub.SubscriptionHandlerDeps{
		SubRepo: subRepository,
		SubService: subService,
	})
	// stat.NewStatHandler(router, &stat.StatHandlerDeps{
	// 	StatRepository: statRepository,
	// }) ? пока под вопросом
	// Middlewares
	// stack := middleware.Chain(
	// 	middleware.Logging,
	// )
	// return stack(router)
	return router
}

func main() {
	// Config
	config.Init()
	logConfig := config.NewLogConfig()
	databaseConfig := config.NewDatabaseConfig()

	// Logger
	customLogger := logger.NewLogger(logConfig)

	app := App(databaseConfig)
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	customLogger.Info().
		Msg("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		customLogger.Error().
			Err(err).
			Msg("Server failed to start")
	}
}
