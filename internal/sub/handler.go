package sub

import "net/http"

type SubscriptionHandlerDeps struct {
	SubRepo *SubscriptionRepository
}

type SubscriptionHandler struct {
	SubRepo *SubscriptionRepository
}

func NewSubscriptionHandler(router *http.ServeMux, deps *SubscriptionHandlerDeps) {
	handler := &SubscriptionHandler{
		SubRepo: deps.SubRepo,
	}

	router.Handle("POST /sub", handler.Create())
	router.Handle("GET /sub/{id}", handler.Read())
	router.Handle("PATCH /sub/{id}", handler.Update())
	router.Handle("DELETE /sub/{id}", handler.Delete())
	router.Handle("GET /sub", handler.List())
}

func (handler *SubscriptionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
func (handler *SubscriptionHandler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
func (handler *SubscriptionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
func (handler *SubscriptionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
func (handler *SubscriptionHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
