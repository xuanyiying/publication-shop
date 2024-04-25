// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xuanyiying/publication-shop/app/payment/service/internal/biz"
	"github.com/xuanyiying/publication-shop/app/payment/service/internal/conf"
	"github.com/xuanyiying/publication-shop/app/payment/service/internal/data"
	"github.com/xuanyiying/publication-shop/app/payment/service/internal/server"
	"github.com/xuanyiying/publication-shop/app/payment/service/internal/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(logger)
	if err != nil {
		return nil, nil, err
	}
	paymentRepo := data.NewPaymentRepo(dataData, logger)
	paymentUseCase := biz.NewPaymentUseCase(paymentRepo, logger)
	paymentService := service.NewPaymentService(paymentUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, paymentService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
