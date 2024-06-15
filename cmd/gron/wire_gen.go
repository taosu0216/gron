// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"gron/internal/biz"
	"gron/internal/conf"
	"gron/internal/data"
	"gron/internal/interfaces"
	"gron/internal/server"
	"gron/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewCache(confData)
	dataData := data.NewData(db, client)
	timerRepo := data.NewGronRepo(dataData)
	timerTaskRepo := data.NewTimerTaskRepo(dataData)
	taskCache := data.NewTaskCache(confData, dataData)
	transaction := data.NewTransaction(dataData)
	migratorUseCase := biz.NewMigratorUseCase(confData, timerRepo, timerTaskRepo, taskCache)
	gronUseCase := biz.NewCreateTimerUseCase(confData, timerRepo, timerTaskRepo, taskCache, transaction, migratorUseCase)
	gronService := service.NewXTimerService(gronUseCase)
	grpcServer := server.NewGRPCServer(confServer, gronService)
	handler := interfaces.NewHandler(gronService)
	httpServer := server.NewHTTPServer(confServer, handler)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}
