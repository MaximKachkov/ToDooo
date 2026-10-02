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

func (h *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...)

	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}
	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		h.logger.Warn("HTTP server launched", zap.String("addr", h.config.Addr))
		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			h.logger.Error("server shutdown by error", zap.String("error", err.Error()))
			ch <- err
		}
	}()
	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("HTTP listen and server error : %w", err)
		}
	case <-ctx.Done():
		h.logger.Warn("context done ")

		SrvContext, cancel := context.WithTimeout(context.Background(), h.config.ShutDownTimeout)
		defer cancel()

		if err := server.Shutdown(SrvContext); err != nil {
			_ = server.Close()

			return fmt.Errorf("server closed : %w", err)
		}
		h.logger.Warn("server stopped")
	}
	return nil
}
