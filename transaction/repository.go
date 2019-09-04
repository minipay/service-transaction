package transaction

import (
	"context"
	"service-transaction/models"
)

// Repository represent the transaction's Repositorys
type Repository interface {
	Fetch(ctx context.Context, offset int, num int, sort string) (res []*models.Transaction, err error)
}
