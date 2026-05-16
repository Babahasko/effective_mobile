package migrate

import (
	"effective_mobile/internal/sub"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)	
	

func Migrate(db *gorm.DB) {
    err := db.AutoMigrate(
        &sub.Subscription{},
    )
    if err != nil {
        log.Fatal().Err(err).Msg("failed to run migrations")
    }
    log.Info().Msg("migrations completed")
}