package app

import (
	"database/sql"
	"time"

	"github.com/RyouRio/belajar-golang-restful-api-2/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:fredrick3246@tcp(localhost:3306)/belajar_golang_rest_api2")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}