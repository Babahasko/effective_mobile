package res

import (
	"net/http"
	"encoding/json"
)

func Json(w http.ResponseWriter,body any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func JsonError(w http.ResponseWriter, message string, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    data, _ := json.Marshal(map[string]string{"error": message})
    w.Write(data)
}