package app

import (
	"database/sql"
	"mini-project/go-crud/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1111@tcp(localhost:3306)/go_crud")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
