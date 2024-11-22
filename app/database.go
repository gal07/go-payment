package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect() (*sql.DB, error) {

	// Value from ENV
	connector := os.Getenv("GP_DATABASE_CONNECTOR")
	db_user := os.Getenv("GP_DATABASE_USER")
	db_pw := os.Getenv("GP_DATABASE_PW")
	db_port := os.Getenv("GP_DATABASE_PORT")
	db_host := os.Getenv("GP_DATABASE_HOST")
	db_name := os.Getenv("GP_DATABASE_NAME")

	db, err := sql.Open(connector, db_user+":"+db_pw+"@tcp("+db_host+":"+db_port+")/"+db_name+"")
	if err != nil {
		return nil, err
	}
	return db, nil
}
