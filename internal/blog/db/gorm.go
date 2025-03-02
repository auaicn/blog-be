package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectGorm: gorm.DB 연결
func ConnectGorm(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("GORM 연결 실패: %w", err)
	}
	return db, nil
}
