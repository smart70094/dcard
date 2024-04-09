package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "10.101.98.105"
	//host     = "127.0.0.1"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "dcard"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))

	if err != nil {
		return nil, err
	}

	return db, nil
}
