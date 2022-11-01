package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"money-management/src/config"
	"money-management/src/pkg/helpers"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func InitConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
	config.Get().PostgreHost, config.Get().PostgrePort, config.Get().PostgreUsername,
	config.Get().PostgrePassword, config.Get().PostgreDbName)
	gorm, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	
	db, err := gorm.DB()
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