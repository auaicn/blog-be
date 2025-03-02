package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL 드라이버
)

// ConnectSQL: sql.DB 연결
func ConnectSQL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("SQL 연결 실패: %w", err)
	}
	return db, nil
}
