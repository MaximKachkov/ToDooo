package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_middleware "github.com/MaximKachkov/ToDooo/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	logger     *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(config Config, log *core_logger.Logger, middleware ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		logger:     log,
		middleware: middleware,
	}
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, router.WithMiddleware()))
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(s.mux, s.middleware...)

	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}
	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		s.logger.Warn("HTTP server launched", zap.String("addr", s.config.Addr))
		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("server shutdown by error", zap.String("error", err.Error()))
			ch <- err
		}
	}()
	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("HTTP listen and server error : %w", err)
		}
	case <-ctx.Done():
		s.logger.Warn("context done ")

		SrvContext, cancel := context.WithTimeout(context.Background(), s.config.ShutDownTimeout)
		defer cancel()

		if err := server.Shutdown(SrvContext); err != nil {
			_ = server.Close()

			return fmt.Errorf("server closed : %w", err)
		}
		s.logger.Warn("server stopped")
	}
	return nil
}
