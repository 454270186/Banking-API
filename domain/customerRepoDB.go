package domain

import (
	"database/sql"
	"log"
	"time"
)

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

type CustomerRepoDB struct {
	DB *sql.DB
}

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", "user=postgres password=2021110003 host=localhost port=5432 dbname=restful")
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

func NewCustomerRepoDB() CustomerRepoDB {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	return CustomerRepoDB{db}
}

func (cdb CustomerRepoDB) FindAll() ([]Customer, error) {
	findAllSql := "Select * from customers"

	rows, err := cdb.DB.Query(findAllSql)
	if err != nil {
		panic(err)
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Birthday, &c.Status)
		if err != nil {
			log.Println("ERROR: while querying customer table: " + err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}