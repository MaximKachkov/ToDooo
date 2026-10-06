package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_pgx_pool "github.com/MaximKachkov/ToDooo/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/MaximKachkov/ToDooo/internal/core/transport/http/middleware"
	core_http_server "github.com/MaximKachkov/ToDooo/internal/core/transport/http/server"
	user_postgres_repository "github.com/MaximKachkov/ToDooo/internal/features/users/repository/postgres"
	users_service "github.com/MaximKachkov/ToDooo/internal/features/users/service"
	users_transport_http "github.com/MaximKachkov/ToDooo/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	logger, err := core_logger.NewLogger(*core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing connection pool")
	pool, err := core_pgx_pool.NewPool(ctx, *core_pgx_pool.ConfigMust())
	if err != nil {
		logger.Fatal("error creating a pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("starting ToDo application!")

	usersRepository := user_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTranposrtHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing HTTP server")

	httpServer := core_http_server.NewHTTPServer(*core_http_server.ConfigMust(), logger,
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic())

	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTranposrtHTTP.Routes()...)
	/*
	   //	apiVersionRouterV2 := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion2, core_http_middleware.Dummy("api v2 middleware"))
	   //	apiVersionRouterV2.RegisterRoutes(usersTranposrtHTTP.Routes()...)
	*/
	httpServer.RegisterAPIRouters(apiVersionRouterV1) //apiVersionRouterV2,

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error ", zap.Error(err))
	}
}
