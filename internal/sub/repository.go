package sub

import (
	"effective_mobile/pkg/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *SubscriptionRepository) Update(sub *Subscription) (*Subscription, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(sub)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
        return nil, gorm.ErrRecordNotFound
    }
	return sub, nil
}

func (repo *SubscriptionRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Subscription{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
	return nil
}

func (repo *SubscriptionRepository) List(filter *SubscriptionFilter) ([]*Subscription, error) {
    var subs []*Subscription
    query := repo.Database.DB.Model(&Subscription{})

    if filter.UserID != nil {
        query = query.Where("user_id = ?", filter.UserID)
    }
    if filter.ServiceName != nil {
        query = query.Where("service_name = ?", filter.ServiceName)
    }
    if filter.Price != nil {
        query = query.Where("price = ?", filter.Price)
    }
    if filter.StartDate != nil {
        query = query.Where("start_date = ?", filter.StartDate)
    }

    result := query.Find(&subs)
    return subs, result.Error
}

func (repo *SubscriptionRepository) Total(filter *SubscriptionFilter) (uint, error) {
    var total uint
    query := repo.Database.DB.Model(&Subscription{}).Select("COALESCE(SUM(price), 0)")

    if filter.UserID != nil {
        query = query.Where("user_id = ?", filter.UserID)
    }
    if filter.ServiceName != nil {
        query = query.Where("service_name = ?", filter.ServiceName)
    }
    if filter.StartDate != nil {
        query = query.Where("start_date >= ?", filter.StartDate)
    }

    result := query.Scan(&total)
    return total, result.Error
}

func (repo *SubscriptionRepository) Exists(serviceName string, userID uuid.UUID) (bool, error) {
    var count int64
    err := repo.Database.Model(&Subscription{}).
        Where("service_name = ? AND user_id = ?", serviceName, userID).
        Count(&count).Error
    return count > 0, err
}