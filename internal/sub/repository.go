package sub

import (
	"effective_mobile/pkg/db"

	"github.com/google/uuid"
)

type SubscriptionRepository struct {
	Database *db.Db
}

func NewSubscriptionRepository(database *db.Db) *SubscriptionRepository {
	return &SubscriptionRepository{
		Database: database,
	}
}

func (repo *SubscriptionRepository) Create(sub *Subscription) (*Subscription, error) {
	result := repo.Database.DB.Create(sub)
	if result.Error != nil {
		return nil, result.Error
	}
	return sub, nil
}

func (repo *SubscriptionRepository) Read(id uint64) (*Subscription, error) {
	var sub *Subscription
	result := repo.Database.DB.First(&sub, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return sub, nil
}

func (repo *SubscriptionRepository) Exists(serviceName string, userID uuid.UUID) (bool, error) {
    var count int64
    err := repo.Database.Model(&Subscription{}).
        Where("service_name = ? AND user_id = ?", serviceName, userID).
        Count(&count).Error
    return count > 0, err
}