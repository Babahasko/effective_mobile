package sub

import "effective_mobile/pkg/db"

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