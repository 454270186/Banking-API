package domain

import (
	"RESTful/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
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

func NewCustomerRepoDB(db *sql.DB) CustomerRepoDB {
	return CustomerRepoDB{db}
}

func (cdb CustomerRepoDB) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "Select * from customers"

	rows, err := cdb.DB.Query(findAllSql)
	if err != nil {
		panic(err)
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		var updatedAt time.Time
		var createAt time.Time
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Birthday, &c.Status, &updatedAt, &createAt)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("Error: while scan customer table")
				return nil, errs.NewNotFoundError("Customers not found")
			} else {
				log.Println("Error: " + err.Error())
				return nil, errs.NewInterError("unexpected database error")
			}
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (cdb CustomerRepoDB) FindById(id string) (*Customer, *errs.AppError) {
	var c Customer
	var updatedAt time.Time
	var createAt time.Time
	byIdSql := "Select * from customers where id = $1"

	row := cdb.DB.QueryRow(byIdSql, id)
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Birthday, &c.Status, &createAt, &updatedAt)
	if err != nil {
		if (err == sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while find customer by id")
			return nil, errs.NewInterError("unexpect database error")
		}
	}

	return &c, nil
}