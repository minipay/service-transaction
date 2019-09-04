package transaction

import (
	"context"
	"service-transaction/models"
)

// Usecase represent the transaction's usecases
type Usecase interface {
	Fetch(ctx context.Context, num int, offset int, sort string) ([]*models.Transaction, error)
}
