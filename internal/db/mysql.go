package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func New(dsn string , driver string) (*DB , error){

	if driver == ""{
		driver = "mysql"
	}

	conn , err := sql.Open(driver , dsn)
	if err != nil {
		return nil, fmt.Errorf("db: open failed: %w", err)
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	if err := conn.Ping() ; err != nil {
		return  nil , fmt.Errorf("db: ping failed: %w", err)
	}

	slog.Info("Database Connection Established")

	return &DB{conn} , nil
}

func (db *DB) Close() error {
	slog.Info("Database Connection Closed")
	return db.DB.Close()
}

func (db *DB) Ping(ctx context.Context) error {
	return db.DB.PingContext(ctx)
}