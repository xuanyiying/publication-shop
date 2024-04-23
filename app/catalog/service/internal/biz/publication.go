package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/publication-shop/pkg/page_token"
	"golang.org/x/sync/singleflight"

	"github.com/go-kratos/kratos/v2/log"
)

type Catalog struct {
	Id              int64
	CatalogCode     string
	CatalogName     string
	ParentCatalogId string
}
type Classic struct {
	ClassicId   int64
	ClassicName string
}
type Language struct {
	LanguageId   int64
	LanguageName string
	LanguageCode string
}
type Author struct {
	AuthorId      int64
	Isbn          string
	PublicationId int64
	Author        string
	AuthorAbout   string
}
type Detail struct {
	DetailId      int64
	Isbn          string
	DetailHtml    string
	DetailImg     string
	PublicationId int64
}
type Image struct {
	ImgId         int64
	ImgUrl        string
	Isbn          int64
	ImgEncode     string
	PublicationId int64
	MainFlag      int32
}
type Publication struct {
	PublicationId   int64
	PublicationName string
	OrgId           string
	PublishedTimes  string
	PrintTimes      string
	Price           float64
	Introduction    string
	Quantity        int32
	WordCount       string
	Isbn            string
	StorageBy       string
	StorageAt       time.Time
	ModifiedBy      string
	CreatedAt       time.Time
	ModifiedAt      time.Time
	Org             Org
	images          []Image
	Authors         []Author
	Stock           Stock
	Catalogs        []Catalog
	Languages       []Language
	Classic         []Classic
	Detail          Detail
}
type Org struct {
	OrgId      int64
	OrgCode    string
	OrgName    string
	OrgAddress string
	ModifiedAt time.Time
	CreatedAt  time.Time
}
type PublicationXCatalog struct {
	PublicationCatalogId int64
	CatalogId            int64
	Isbn                 string
	CatalogName          string
	PublicationId        int64
	ParentCatalogId      int64
}
type PublicationXClassic struct {
	PublicationClassicId int64
	ClassicId            int64
	Isbn                 string
	ClassicName          string
	PublicationId        int64
}
type PublicationXLanguage struct {
	PublicationLanguageId int64
	LanguageId            int64
	Isbn                  string
	LanguageName          string
	PublicationId         int64
}

type Stock struct {
	StockId       int64
	Isbn          string
	Quantity      int64
	PublicationId int64
}

type PublicationRepo interface {
	CreatePublication(ctx context.Context, c *Publication) (*Publication, error)
	UpdatePublication(ctx context.Context, c *Publication) (*Publication, error)
	GetPublication(ctx context.Context, id int64) (*Publication, error)
	ListPublication(ctx context.Context, pageNum, pageSize int64) ([]*Publication, error)
	ListPublicationNext(ctx context.Context, start, end int32) ([]*Publication, error)
	Count(ctx context.Context) (int, error)
}

type PublicationUseCase struct {
	repo      PublicationRepo
	log       *log.Helper
	pageToken page_token.ProcessPageTokens
	sg        singleflight.Group
}

func NewPublication(repo PublicationRepo, logger log.Logger) *PublicationUseCase {
	return &PublicationUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/Publication")), pageToken: page_token.NewTokenGenerate()}
}

func (uc *PublicationUseCase) Create(ctx context.Context, u *Publication) (*Publication, error) {
	return uc.repo.CreatePublication(ctx, u)
}

func (uc *PublicationUseCase) Get(ctx context.Context, id int64) (*Publication, error) {
	return uc.repo.GetPublication(ctx, id)
}

func (uc *PublicationUseCase) Update(ctx context.Context, u *Publication) (*Publication, error) {
	return uc.repo.UpdatePublication(ctx, u)
}

func (uc *PublicationUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Publication, error) {
	return uc.repo.ListPublication(ctx, pageNum, pageSize)
}

func (uc *PublicationUseCase) ListNext(ctx context.Context, pageSize int32, pageToken string) ([]*Publication, string, error) {
	total, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, "", err
	}

	start, end, nextToken, err := uc.pageToken.ProcessPageTokens(total, pageSize, pageToken)
	if err != nil {
		return nil, "", err
	}
	// single flight
	data, err, _ := uc.sg.Do(fmt.Sprintf("list_next_%d_%d", start, end), func() (interface{}, error) {
		return uc.repo.ListPublicationNext(ctx, start, end)
	})

	return data.([]*Publication), nextToken, err
}
