package sub

import (
	"time"

	"github.com/google/uuid"
)

type MonthYear struct {
	time.Time
}

type Subscription struct {
	ID          uint    `gorm:"primarykey"`
	ServiceName string    `json:"service_name" gorm:"column:service_name;uniqueIndex:unique_user_service"`
	UserID      uuid.UUID `json:"user_id"      gorm:"column:user_id;type:uuid;uniqueIndex:unique_user_service"`
	Price       uint      `json:"price"`
	StartDate   time.Time `json:"start_date"   gorm:"column:start_date;type:date"`
}
