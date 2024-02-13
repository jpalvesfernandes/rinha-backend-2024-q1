package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	dsn := "host=db user=rinha password=123 dbname=rinha port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)  // Set maximum idle connections
	sqlDB.SetMaxOpenConns(500) // Set maximum open connections

	return db, nil
}
