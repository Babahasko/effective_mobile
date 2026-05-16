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
