package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"money-management/src/config"
	"money-management/src/pkg/helpers"
	"time"

	_ "github.com/lib/pq"
)

func InitConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	config.Get().PostgreHost, config.Get().PostgrePort, config.Get().PostgreUsername,
	config.Get().PostgrePassword, config.Get().PostgreDbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.Get().PostgreMaxConnection)
	db.SetMaxIdleConns(config.Get().PostgreMaxIdleConnection)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) {
	err := db.Close()
	fmt.Println("Close db connection")
	if err != nil {
		helper.Logger.Error(err.Error())
	}
}