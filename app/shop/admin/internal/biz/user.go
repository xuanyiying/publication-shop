package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	usV1 "github.com/xuanyiying/publication-shop/api/user/service/v1"
)

type User struct {
	Id int64
}

type UserRepo interface {
}

type UserUseCase struct {
	repo UserRepo
	us   usV1.UserClient
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger, us usV1.UserClient) *UserUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		repo: repo,
		us:   us,
		log:  log,
	}
}
