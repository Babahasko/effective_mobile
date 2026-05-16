package logger

import (
	"effective_mobile/config"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(cfg *config.LogConfig) {
    zerolog.SetGlobalLevel(zerolog.Level(cfg.Level))

    writers := []io.Writer{os.Stdout}
    if cfg.FilePath != "" {
        file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Warn().Err(err).Msg("failed to open log file")
        } else {
            writers = append(writers, file)
        }
    }

    multi := zerolog.MultiLevelWriter(writers...)
    log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}