package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func NewPostgres(
	host string,
	port int,
	user string,
	password string,
	dbname string,
	maxOpenConns int,
	maxIdleConns int,
	maxIdleTime time.Duration,
) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
