module github.com/go-kratos/publication-shop

go 1.21

toolchain go1.22.2

require (
	entgo.io/ent v0.13.1
	github.com/Shopify/sarama v1.29.1
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220223091357-d6896127b137
	github.com/go-kratos/kratos/v2 v2.2.0
	github.com/go-kratos/swagger-api v0.1.6
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/golang-jwt/jwt/v4 v4.2.0
	github.com/golang/protobuf v1.5.4
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/consul/api v1.11.0
	github.com/stretchr/testify v1.8.2
	go.mongodb.org/mongo-driver v1.5.1
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/sdk v1.3.0
	golang.org/x/crypto v0.19.0
	golang.org/x/sync v0.6.0
	google.golang.org/genproto v0.0.0-20220126215142-9970aeb2e350
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.33.0
	gorm.io/driver/mysql v1.0.6
	gorm.io/gorm v1.21.9
)

require (
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
)
