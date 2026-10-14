package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_http_server "github.com/MaximKachkov/ToDooo/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(ctx context.Context, id *int, from *time.Time, to *time.Time) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(stat StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticsService: stat,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
