package sub

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SubscriptionCreateRequest struct {
	ServiceName string    `json:"service_name" validate:"required"`
	Price       uint      `json:"price" validate:"required"`
	UserId      uuid.UUID `json:"user_id" validate:"required,uuid"`
	StartDate   string    `json:"start_date" validate:"required,month_year"`
}

type SubscriptionUpdateRequest struct {
	ServiceName string    `json:"service_name"`
	Price       uint      `json:"price"`
	UserId      uuid.UUID `json:"user_id"`
	StartDate   string    `json:"start_date" validate:"month_year"`
}

func (r *SubscriptionCreateRequest) ToModel() (*Subscription, error) {
	t, err := time.Parse("01-2006", r.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start_date: %w", err)
	}
	return &Subscription{
		ServiceName: r.ServiceName,
		Price:       r.Price,
		UserID:      r.UserId,
		StartDate:   t,
	}, nil
}

func (r *SubscriptionUpdateRequest) ToModel() (*Subscription, error) {
	t, err := time.Parse("01-2006", r.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start_date: %w", err)
	}
	return &Subscription{
		ServiceName: r.ServiceName,
		Price:       r.Price,
		UserID:      r.UserId,
		StartDate:   t,
	}, nil
}

type SubscriptionResponse struct {
    ID          uint      `json:"id"`
    ServiceName string    `json:"service_name"`
    UserID      uuid.UUID `json:"user_id"`
    Price       uint      `json:"price"`
    StartDate   string    `json:"start_date"`
}

type TotalResponse struct {
    Total       uint       `json:"total"`
    UserID      *uuid.UUID `json:"user_id,omitempty"`
    ServiceName *string    `json:"service_name,omitempty"`
    StartDate   *string    `json:"start_date,omitempty"`
}

func NewSubscriptionResponse(sub *Subscription) *SubscriptionResponse {
    return &SubscriptionResponse{
        ID:          sub.ID,
        ServiceName: sub.ServiceName,
        UserID:      sub.UserID,
        Price:       sub.Price,
        StartDate:   sub.StartDate.Format("01-2006"),
    }
}
