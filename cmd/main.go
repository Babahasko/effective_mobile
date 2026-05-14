package main

import (
	"effective_mobile/config"
	"effective_mobile/logger"
	"net/http"
)
func App() http.Handler{
	
	// db := db.NewDB(conf)
	router := http.NewServeMux()

	//Repositories
	// subRepository := link.NewLinkRepository(db)
	// statRepository := stat.NewStatRepository(db) // пока под вопросом

	// Handler
	// sub.NewSubHandler(router, link.LinkHandlerDeps{
	// 	LinkRepository: subRepository,
	// 	Config:         conf,
	// 	EventBus:       eventBus, ? nтоже под вопросом
	// })
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

	// Logger
	customLogger := logger.NewLogger(logConfig)

	app := App()
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
