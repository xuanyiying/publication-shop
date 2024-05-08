package service

import (
	v1 "github.com/xuanyiying/publication-shop/api/transaction/service/v1"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/biz"
)

// TransactionService is a greeter service.
type TransactionService struct {
	v1.UnimplementedTransactionServer

	uc *biz.TransactionUsecase
}

// NewTransactionService new a greeter service.
func NewTransactionService(uc *biz.TransactionUsecase) *TransactionService {
	return &TransactionService{uc: uc}
}
