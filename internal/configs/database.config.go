package configs

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	DB *sqlx.DB
}

func NewDatabaseConnection() *DatabaseConnection {
	dbConnectionConfig := &DatabaseConnection{}
	dbConnectionConfig.connect()

	return dbConnectionConfig
}

func (conn *DatabaseConnection) connect() {
	appConfig := NewConfig()
	dbConnectionString := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=5432",
		appConfig.DBUsername,
		appConfig.DBName,
		appConfig.DBEnableSSL,
		appConfig.DBPassword, appConfig.DBHost)

	db, err := sqlx.Connect("postgres", dbConnectionString)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("failed to ping database: %w", err))
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	conn.DB = db
}
