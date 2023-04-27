package domain

import (
	"RESTful/global"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

var dbUser string
var dbPassword string
var dbHost string
var dbPort int
var dbName string

func initDB() {
	dbUser = global.Settings.PostgresInfo.DBUser
	dbPassword = global.Settings.PostgresInfo.DBPassword
	dbHost = global.Settings.PostgresInfo.DBHost
	dbPort = global.Settings.PostgresInfo.DBPort
	dbName = global.Settings.PostgresInfo.DBName
}

func NewDB() (*sql.DB, error) {
	initDB()
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	return db, nil
}
