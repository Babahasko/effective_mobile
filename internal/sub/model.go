package sub

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MonthYear struct {
	time.Time
}

type Subscription struct {
    ID          uint      `gorm:"primarykey"`
    ServiceName string    `json:"service_name" gorm:"column:service_name"`
    Price       int       `json:"price"`
    UserID      uuid.UUID `json:"user_id"      gorm:"column:user_id;type:uuid"`
    StartDate   time.Time `json:"start_date"   gorm:"column:start_date;type:date"`
}

type SubscriptionDTO struct {
    ServiceName string `json:"service_name"`
    Price       int    `json:"price"`
    UserID      string `json:"user_id"`
    StartDate   string `json:"start_date"` // "07-2025"
}

func (dto *SubscriptionDTO) ToModel() (*Subscription, error) {
    t, err := time.Parse("01-2006", dto.StartDate)
    if err != nil {
        return nil, fmt.Errorf("invalid start_date format, expected MM-YYYY")
    }
    uid, err := uuid.Parse(dto.UserID)
    if err != nil {
        return nil, fmt.Errorf("invalid user_id format")
    }
    return &Subscription{
        ServiceName: dto.ServiceName,
        Price:       dto.Price,
        UserID:      uid,
        StartDate:   t,
    }, nil
}
