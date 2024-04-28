package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/conf"
	"github.com/xuanyiying/publication-shop/app/book/service/internal/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(data.NewBookRepo, NewData, NewDB)

type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "order-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatal(err)
	}
	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "book-service/data"))

	d := &Data{
		db:  db,
		log: log,
	}
	return d, func() {

	}, nil
}