package postgre

import (
	"context"
	"database/sql"
	"money-management/src/config"
	"time"
)

func InitConnection() (*sql.DB, error) {
	db, err := sql.Open("postgre", config.Get().PostgreHost)
	if err != nil {
		//todo add logger
		return nil, err
	}

	db.SetMaxOpenConns(config.Get().PostgreMaxConnection)
	db.SetMaxIdleConns(config.Get().PostgreMaxIdleConnection)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		//todo add logger
		return nil, err
	}

	//todo add logger
	return db, nil
}

func CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		//todo add logger
	}
}