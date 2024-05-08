package data

import (
	"context"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/transaction"
	"github.com/xuanyiying/publication-shop/app/transaction/internal/data/ent/txitem"

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
	po, err := r.data.db.Transaction.
		Create().
		SetUserID(g.UserId).
		SetQuantity(g.Quantity).
		SetTxAmount(g.TxAmount).
		SetTxDate(g.TxDate).
		SetTxNo(g.TxNo).
		SetTxStatus(transaction.TxStatus(g.TxStatus)).
		SetTxType(transaction.TxType(g.TxType)).
		SetPaymentID(g.PaymentId).
		SetID(g.TxId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	r.AddTxItems(ctx, g.TxItems, po)
	return &biz.Transaction{TxId: po.ID, UserId: po.UserID, TxNo: po.TxNo, TxStatus: string(po.TxStatus), TxType: string(po.TxType), TxDate: po.TxDate, TxAmount: po.TxAmount, PaymentId: po.PaymentID}, nil
}

func (r *txRepo) AddTxItems(ctx context.Context, items []*biz.TxItem, po *ent.Transaction) {
	for _, item := range items {
		r.data.db.TxItem.Create().
			SetTxID(po.ID).
			SetTxType(txitem.TxType(item.TxType)).
			SetIsbn(item.Isbn).
			SetPrice(item.Price).
			SetAuthor(item.Author).
			SetQuantity(item.Quantity).
			SetImageURL(item.ImageUrl).
			SetBookID(item.BookId).
			SetPublisherID(item.PublisherId).
			SetTitle(item.Title).
			SetID(item.TxItemId).
			Save(ctx)
		r.log.Infof("create tx item: %d", item.TxItemId)
	}
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
