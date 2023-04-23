package domain

import (
	"RESTful/errs"
	"RESTful/logger"
	"database/sql"
	"log"
)

type AccountRepoDB struct {
	DB *sql.DB
}

func NewAccountRepoDB(db *sql.DB) AccountRepoDB {
	return AccountRepoDB{db}
}

func (adb AccountRepoDB) Save(a Account) (*Account, *errs.AppError) {
	insertSql := `
		INSERT INTO accounts (customer_id, opening_date, account_type, amount, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING account_id
	`

	err := adb.DB.QueryRow(insertSql, a.CustomerID, a.OpeningDate, a.Type, a.Amount, a.Status).Scan(&a.Id)
	if err != nil {
		logger.Error("Error while insert a new account")
		log.Println(err.Error())
		return nil, errs.NewInterError("error while insert a new account")
	}

	return &a, nil
}