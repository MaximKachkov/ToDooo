package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/MaximKachkov/ToDooo/docs"
	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_pgx_pool "github.com/MaximKachkov/ToDooo/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/MaximKachkov/ToDooo/internal/core/transport/http/middleware"
	core_http_server "github.com/MaximKachkov/ToDooo/internal/core/transport/http/server"
	statistics_repository_postgres "github.com/MaximKachkov/ToDooo/internal/features/statistics/repository/postgres"
	statistics_service "github.com/MaximKachkov/ToDooo/internal/features/statistics/service"
	statistics_transport_http "github.com/MaximKachkov/ToDooo/internal/features/statistics/transport/http"
	tasks_postgres_repository "github.com/MaximKachkov/ToDooo/internal/features/tasks/repository/postgres"
	tasks_service "github.com/MaximKachkov/ToDooo/internal/features/tasks/service"
	tasks_transport_http "github.com/MaximKachkov/ToDooo/internal/features/tasks/transport/http"
	user_postgres_repository "github.com/MaximKachkov/ToDooo/internal/features/users/repository/postgres"
	users_service "github.com/MaximKachkov/ToDooo/internal/features/users/service"
	users_transport_http "github.com/MaximKachkov/ToDooo/internal/features/users/transport/http"
	web_fs_repository "github.com/MaximKachkov/ToDooo/internal/features/web/repository/file_system"
	web_service "github.com/MaximKachkov/ToDooo/internal/features/web/service"
	web_transport_http "github.com/MaximKachkov/ToDooo/internal/features/web/transport/http"
	"go.uber.org/zap"
)

var (
	timeZone = time.UTC
)

// @title 	Golang ToDo API
// @version 	1.0
// @description 	ToDo app REST-API scheme
// @host 	127.0.0.1:5050
// @BasePath 	/api/v1
func main() {
	time.Local = timeZone
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

	tasksRepository := tasks_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	statisticsRepository := statistics_repository_postgres.NewTasksRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticsHTTPHandler(statisticsService)

	webRepository := web_fs_repository.NewWebRepository()
	webService := web_service.NewWebService(webRepository)
	webTransportHTTP := web_transport_http.NewWebHTTPHandler(webService)

	logger.Debug("initializing HTTP server")

	httpServer := core_http_server.NewHTTPServer(*core_http_server.ConfigMust(), logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic())

	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTranposrtHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(tasksTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(statisticsTransportHTTP.Routes()...)

	/*
	   //	apiVersionRouterV2 := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion2, core_http_middleware.Dummy("api v2 middleware"))
	   //	apiVersionRouterV2.RegisterRoutes(usersTranposrtHTTP.Routes()...)
	*/
	httpServer.RegisterAPIRouters(apiVersionRouterV1) //apiVersionRouterV2,
	httpServer.RegisterSwagger()
	httpServer.RegisterRoutes(webTransportHTTP.Routes()...)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error ", zap.Error(err))
	}
}
