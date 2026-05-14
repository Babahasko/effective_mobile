package sub

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type MonthYear struct {
	time.Time
}

const monthYearLayout = "01-2006"

func (m *MonthYear) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	t, err := time.Parse(monthYearLayout, s)
	if err != nil {
		return err
	}
	m.Time = t
	return nil
}

func (m MonthYear) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.Time.Format(monthYearLayout) + `"`), nil
}

type Subscription struct {
	ID          uint      `gorm:"primarykey"`
	ServiceName string    `json:"service_name" gorm:"column:service_name"`
	Price       int       `json:"price"`
	UserID      uuid.UUID `json:"user_id"      gorm:"column:user_id;type:uuid"`
	StartDate   MonthYear `json:"start_date"   gorm:"column:start_date"`
}
