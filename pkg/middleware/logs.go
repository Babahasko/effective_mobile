package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		body := bytes.TrimSpace(wrapper.Body.Bytes())
		event := log.Info().
            Int("status", wrapper.StatusCode).
            Str("method", r.Method).
            Str("path", r.URL.Path).
            Dur("duration", time.Since(start))
        if json.Valid(body) {
            event = event.RawJSON("response", body)
        }
        event.Send()
	})
}
