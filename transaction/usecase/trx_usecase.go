package usecase

import (
	"context"
	"time"

	"service-transaction/models"
	"service-transaction/transaction"
)

type transactionUsecase struct {
	transactionRepo transaction.Repository
	contextTimeout  time.Duration
}

// NewTransactionUsecase will create new an transactionUsecase object representation of article.Usecase interface
func NewTransactionUsecase(tr transaction.Repository, timeout time.Duration) transaction.Usecase {
	return &transactionUsecase{
		transactionRepo: tr,
		contextTimeout:  timeout,
	}
}

func (tu *transactionUsecase) Fetch(c context.Context, num int, offset int, sort string) ([]*models.Transaction, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	listTransaction, err := tu.transactionRepo.Fetch(ctx, num, offset, sort)
	if err != nil {
		return nil, err
	}

	return listTransaction, nil
}
