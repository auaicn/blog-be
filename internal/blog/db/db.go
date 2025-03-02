package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBManager 구조체: sql.DB와 gorm.DB를 모두 관리
type DBManager struct {
	SQL  *sql.DB
	Gorm *gorm.DB
}

var DB *DBManager // 전역 변수로 DBManager 선언

// ConnectBoth: SQL과 GORM 둘 다 연결
func ConnectBoth(dsn string) error {
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("SQL 연결 실패: %w", err)
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("GORM 연결 실패: %w", err)
	}

	DB = &DBManager{SQL: sqlDB, Gorm: gormDB} // 전역 변수에 할당
	log.Println("✅ DB 연결 성공")
	return nil
}
