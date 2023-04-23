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

func (adb AccountRepoDB) FindBy(id string) (*Account, *errs.AppError) {
	selectSql := `
		SELECT * FROM accounts WHERE account_id = $1
	`

	var a Account
	log.Println("id = " + id)
	err := adb.DB.QueryRow(selectSql, id).Scan(&a.Id, &a.CustomerID, &a.OpeningDate, &a.Type, &a.Amount, &a.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("no rows")
			return nil, errs.NewInterError("Unexpect database error")
		}
		logger.Error("Error while find account by id")
		log.Println(err.Error())
		return nil, errs.NewInterError("Unexpect database error")
	}

	return &a, nil
}

func (adb AccountRepoDB) SaveTransaction(ts Transaction) (*Transaction, *errs.AppError) {
	insertSql := `
		INSERT INTO transaction (account_id, amount, transaction_type, transaction_date)
		VALUES ($1, $2, $3, $4)
		RETURNING transaction_id
	`

	// start transaction
	tx, err := adb.DB.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account" + err.Error())
		return nil, errs.NewInterError("Unexpect database error")
	}

	err = tx.QueryRow(insertSql, ts.AccountID, ts.Amount, ts.TransactionType, ts.TransactionDate).Scan(&ts.TransactionID)
	if err != nil {
		logger.Error("Error while insert a new transaction" + err.Error())
		return nil, errs.NewInterError("Unexpect database error")
	}

	// update account
	if ts.IsWithdrawal() {
		updateSql := "UPDATE accounts SET amount = amount - $1 WHERE account_id = $2"
		_, err = tx.Exec(updateSql, ts.Amount, ts.AccountID);
	} else {
		updateSql := "UPDATE accounts SET amount = amount + $1 WHERE account_id = $2"
		_, err = tx.Exec(updateSql, ts.Amount, ts.AccountID)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction")
		return nil, errs.NewInterError("Unexpect database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction")
		return nil, errs.NewInterError("Unexpect database error")
	}

	account, appErr := adb.FindBy(ts.AccountID)
	if appErr != nil {
		return nil, appErr
	}

	ts.Amount = account.Amount
	return &ts, nil
}