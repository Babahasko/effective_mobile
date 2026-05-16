package logger

import (
	"effective_mobile/config"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(config *config.LogConfig) {
	zerolog.SetGlobalLevel(zerolog.Level(config.Level))
	 consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
    
    writers := []io.Writer{consoleWriter}
	if config.FilePath != "" {
		file, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Warn().Err(err).Msg("failed to open log file")
		} else {
			writers = append(writers, file)  // file — это обычный io.Writer, пишет сырой JSON
		}
	}
	multi := zerolog.MultiLevelWriter(writers...)
    log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}