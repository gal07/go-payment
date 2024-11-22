package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect() (*sql.DB, error) {

	// Value from ENV
	connector := os.Getenv("DATABASE_CONNECTOR")
	db_user := os.Getenv("DATABASE_USER")
	db_pw := os.Getenv("DATABASE_PW")
	db_port := os.Getenv("DATABASE_PORT")
	db_host := os.Getenv("DATABASE_HOST")
	db_name := os.Getenv("DATABASE_NAME")

	db, err := sql.Open(connector, db_user+":"+db_pw+"@tcp("+db_host+":"+db_port+")/"+db_name+"")
	if err != nil {
		return nil, err
	}
	return db, nil
}
