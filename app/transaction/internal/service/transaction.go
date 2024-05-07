package service

import (
	"context"

	v1 "github.com/xuanyiying/publication-shop/api/transaction/service/v1"
)

type TransactionService struct {
	v1.UnimplementedTransactionServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) CreateTx(ctx context.Context, req *v1.CreateTxReq) (*v1.CreateTxResp, error) {
	return &v1.CreateTxResp{}, nil
}
func (s *TransactionService) UpdateTx(ctx context.Context, req *v1.UpdateTxReq) (*v1.UpdateTxResp, error) {
	return &v1.UpdateTxResp{}, nil
}
func (s *TransactionService) DeleteTx(ctx context.Context, req *v1.DeleteTxReq) (*v1.DeleteTxResp, error) {
	return &v1.DeleteTxResp{}, nil
}
func (s *TransactionService) GetTx(ctx context.Context, req *v1.GetTxReq) (*v1.TxResp, error) {
	return &v1.TxResp{}, nil
}
func (s *TransactionService) ListTx(ctx context.Context, req *v1.ListTxReq) (*v1.ListTxResp, error) {
	return &v1.ListTxResp{}, nil
}
func (s *TransactionService) GetTxByTxNo(ctx context.Context, req *v1.GetTxReq) (*v1.TxResp, error) {
	return &v1.TxResp{}, nil
}
func (s *TransactionService) ListTxByTxType(ctx context.Context, req *v1.ListTxReq) (*v1.ListTxResp, error) {
	return &v1.ListTxResp{}, nil
}
func (s *TransactionService) ListTxByUserId(ctx context.Context, req *v1.ListTxReq) (*v1.ListTxResp, error) {
	return &v1.ListTxResp{}, nil
}
func (s *TransactionService) ListTxByPaymentId(ctx context.Context, req *v1.ListTxReq) (*v1.ListTxResp, error) {
	return &v1.ListTxResp{}, nil
}
func (s *TransactionService) ListTxByTxStatus(ctx context.Context, req *v1.ListTxReq) (*v1.ListTxResp, error) {
	return &v1.ListTxResp{}, nil
}
func (s *TransactionService) ListTxByTxDateRange(ctx context.Context, req *v1.ListTxDateRangeReq) (*v1.ListTxResp, error) {
	return &v1.ListTxResp{}, nil
}
