package service

import (
	"context"

	"github.com/publication-shop/api/payment/service/v1"
)

func (s *PaymentService) PaymentAuth(ctx context.Context, req *v1.PaymentAuthReq) (reply *v1.PaymentAuthReply, err error) {
	return &v1.PaymentAuthReply{}, err
}
