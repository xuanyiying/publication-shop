package data

import (
	"context"

	"github.com/xuanyiying/publication-shop/app/transaction/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type txRepo struct {
	data *Data
	log  *log.Helper
}

// NewTransactionRepo .
func NewTransactionRepo(data *Data, logger log.Logger) biz.TransactionRepo {
	return &txRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *txRepo) Save(ctx context.Context, g *biz.Transaction) (*biz.Transaction, error) {
	return g, nil
}

func (r *txRepo) Update(ctx context.Context, g *biz.Transaction) (*biz.Transaction, error) {
	return g, nil
}

func (r *txRepo) FindByID(context.Context, int64) (*biz.Transaction, error) {
	return nil, nil
}

func (r *txRepo) ListByHello(context.Context, string) ([]*biz.Transaction, error) {
	return nil, nil
}

func (r *txRepo) ListAll(context.Context) ([]*biz.Transaction, error) {
	return nil, nil
}
