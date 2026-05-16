package sub

import (
	"effective_mobile/pkg/req"
	"effective_mobile/pkg/res"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type SubscriptionHandlerDeps struct {
	SubRepo *SubscriptionRepository
	SubService *SubscriptionService
}

type SubscriptionHandler struct {
	SubRepo *SubscriptionRepository
	SubService *SubscriptionService
}

func NewSubscriptionHandler(router *http.ServeMux, deps *SubscriptionHandlerDeps) {
	handler := &SubscriptionHandler{
		SubRepo: deps.SubRepo,
		SubService: deps.SubService,
	}

	router.Handle("POST /sub", handler.Create())
	router.Handle("GET /sub/{id}", handler.Read())
	router.Handle("PATCH /sub/{id}", handler.Update())
	router.Handle("DELETE /sub/{id}", handler.Delete())
	router.Handle("GET /sub", handler.List())
}

func (handler *SubscriptionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SubscriptionCreateRequest](r)
		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		model, err := body.ToModel()
		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
            return
		}
		createdSubscription, err := handler.SubService.Create(model)
		if err != nil {
			res.JsonError(w, err.Error(), http.StatusConflict)
			return
		}
		res.Json(w, createdSubscription, http.StatusOK)
	}
}
func (handler *SubscriptionHandler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		sub, err := handler.SubRepo.Read(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				res.JsonError(w, ErrSubNotFound.Error(), http.StatusNotFound)
				return
			}
			res.JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, sub, http.StatusOK)
	}
}
func (handler *SubscriptionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		parsedid, err := strconv.ParseUint(idString, 10, 64)
		id := uint(parsedid)
		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := req.HandleBody[SubscriptionUpdateRequest](r)
		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		model, err := body.ToModel()
		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		sub, err := handler.SubRepo.Update(&Subscription{
			ID: id,
			ServiceName: model.ServiceName,
			UserID: model.UserID,
			Price: model.Price,
			StartDate: model.StartDate,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
                res.JsonError(w, ErrSubNotFound.Error(), http.StatusNotFound)
                return
            }
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, sub, http.StatusOK)
	}
}
func (handler *SubscriptionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		parsedid, err := strconv.ParseUint(idString, 10, 64)
		id := uint(parsedid)
		if err != nil {
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := handler.SubRepo.Delete(id); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
                res.JsonError(w, ErrSubNotFound.Error(), http.StatusNotFound)
                return
            }
			res.JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, fmt.Sprintf("subscription id: %v deleted", id), http.StatusOK)
	}
}
func (handler *SubscriptionHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        filter, err := FilterFromQuery(r)
        if err != nil {
            res.JsonError(w, err.Error(), http.StatusBadRequest)
            return
        }
        subs, err := handler.SubRepo.List(filter)
        if err != nil {
            res.JsonError(w, err.Error(), http.StatusInternalServerError)
            return
        }
        res.Json(w, subs, http.StatusOK)
    }
}
