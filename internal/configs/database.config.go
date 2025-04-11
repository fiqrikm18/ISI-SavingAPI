package configs

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DatabaseConnection struct {
	DB *sqlx.DB
}

func (conn *DatabaseConnection) Connect() {
	dbConnectionString := fmt.Sprintf("")
	db, err := sqlx.Connect("postgres", dbConnectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	conn.DB = db
}
