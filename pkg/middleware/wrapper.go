package middleware

import (
	"bytes"
	"net/http"
)

type WrapperWriter struct {
    http.ResponseWriter
    StatusCode int
    Body       bytes.Buffer
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func (w *WrapperWriter) Write(b []byte) (int, error) {
    w.Body.Write(b)
    return w.ResponseWriter.Write(b)
}