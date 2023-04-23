package domain

import (
	"RESTful/errs"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type CustomerRepoDB struct {
	DB *sql.DB
}

func NewCustomerRepoDB(db *sql.DB) CustomerRepoDB {
	return CustomerRepoDB{db}
}

func (cdb CustomerRepoDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSql string
	if status == "active" {
		findAllSql = "Select * from customers where status = true"
	} else if status == "inactive" {
		findAllSql = "Select * from customers where status = false"
	} else if status == "" {
		findAllSql = "Select * from customers"
	} else {
		log.Println("invalid status")
		return nil, errs.NewNotFoundError("invalid status")
	}

	rows, err := cdb.DB.Query(findAllSql)
	if err != nil {
		panic(err)
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Birthday, &c.Status)
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
	byIdSql := "Select * from customers where id = $1"

	row := cdb.DB.QueryRow(byIdSql, id)
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Birthday, &c.Status)
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