package db

import (
	"effective_mobile/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct{
	*gorm.DB
}

func NewDB(conf *config.DatabaseConfig) *Db {
	db, err := gorm.Open(postgres.Open(conf.Url),  &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}