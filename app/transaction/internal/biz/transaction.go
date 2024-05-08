package biz

import (
	"context"
	"time"

	_ "github.com/xuanyiying/publication-shop/api/transaction/service/v1"

	_ "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// /ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Transaction is a Transaction model.
type Transaction struct {
	TxId      int64
	TxNo      string
	TxType    string
	UserId    int64
	Quantity  int
	TxStatus  string
	TxDate    time.Time
	TxAmount  float64
	PaymentId int64
	TxItems   []*TxItem
}
type TxItem struct {
	TxItemId    int64
	TxType      string
	TxId        int64
	BookId      int64
	Quantity    int
	Price       float64
	Isbn        string
	Title       string
	Author      string
	PublisherId int64
	ImageUrl    string
}

type Orders struct {
	OrderId          int64
	OrderNo          string
	TxId             int64
	OrderStatus      string
	DeliveredAddress string
	ShippingCost     float64
	TotalAmount      float64
	PlacedUserId     int64
	PlacedAt         time.Time
	ShippedAddress   string
	ShippedAt        time.Time
	PaymentId        int64
}

// TransactionRepo is a Greater repo.
type TransactionRepo interface {
	Save(context.Context, *Transaction) (*Transaction, error)
	Update(context.Context, *Transaction) (*Transaction, error)
	FindByID(context.Context, int64) (*Transaction, error)
	ListByHello(context.Context, string) ([]*Transaction, error)
	ListAll(context.Context) ([]*Transaction, error)
}

// TransactionUsecase is a Transaction usecase.
type TransactionUsecase struct {
	repo TransactionRepo
	log  *log.Helper
}

// NewTransactionUsecase new a Transaction usecase.
func NewTransactionUsecase(repo TransactionRepo, logger log.Logger) *TransactionUsecase {
	return &TransactionUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateTransaction creates a Transaction, and returns the new Transaction.
func (uc *TransactionUsecase) CreateTransaction(ctx context.Context, g *Transaction) (*Transaction, error) {
	uc.log.WithContext(ctx).Infof("CreateTransaction: %v", g.TxId)
	return uc.repo.Save(ctx, g)
}
