package repository

import (
	"context"
	"database/sql"
	"log"
	"service-transaction/utils/errorcollector"

	"service-transaction/models"
	"service-transaction/transaction"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlTransactionRepository struct {
	Conn *sql.DB
}

// NewMysqlTransactionRepository will create an object that represent the article.Repository interface
func NewMysqlTransactionRepository(Conn *sql.DB) transaction.Repository {
	return &mysqlTransactionRepository{Conn}
}

func (m *mysqlTransactionRepository) Fetch(ctx context.Context, num int, offset int, sort string) ([]*models.Transaction, error) {
	res := make([]*models.Transaction, 0)
	if num == 0 {
		num = 20
	}
	if sort == "" {
		sort = "ASC"
	}

	q := "select unique_id_transaction, type, amount, description, created_at from transactions order by id " + sort + " limit ? offset ?"
	rows, err := m.Conn.QueryContext(ctx, q, num, offset)
	if err != nil {
		log.Println(err)
		errorcollector.WriteLog(err.Error())
		return res, err
	}
	defer func() {
		rows.Close()
	}()
	for rows.Next() {
		t := new(models.Transaction)
		err = rows.Scan(
			&t.TransactionID,
			&t.Type,
			&t.Amount,
			&t.Desc,
			&t.CreatedAt,
		)
		if err != nil {
			log.Println(err)
			errorcollector.WriteLog(err.Error())
			return res, err
		}
		res = append(res, t)
	}
	return res, nil
}
