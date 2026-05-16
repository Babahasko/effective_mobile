package sub

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type SubscriptionFilter struct {
	UserID      *uuid.UUID `json:"user_id"`
	ServiceName *string    `json:"service_name"`
	Price       *uint      `json:"price"`
	StartDate   *time.Time `json:"start_date"`
}

func FilterFromQuery(r *http.Request) (*SubscriptionFilter, error) {
    filter := &SubscriptionFilter{}

    if v := r.URL.Query().Get("user_id"); v != "" {
        id, err := uuid.Parse(v)
        if err != nil {
            return nil, fmt.Errorf("invalid user_id")
        }
        filter.UserID = &id
    }
    if v := r.URL.Query().Get("service_name"); v != "" {
        filter.ServiceName = &v
    }
    if v := r.URL.Query().Get("price"); v != "" {
        parsed, err := strconv.ParseUint(v, 10, 64)
        if err != nil {
            return nil, fmt.Errorf("invalid price")
        }
        price := uint(parsed)
        filter.Price = &price
    }
    if v := r.URL.Query().Get("start_date"); v != "" {
        date, err := time.Parse("01-2006", v)
        if err != nil {
            return nil, fmt.Errorf("invalid start_date, use format MM-YYYY")
        }
        filter.StartDate = &date
    }

    return filter, nil
}
