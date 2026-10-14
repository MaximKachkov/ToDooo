package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MaximKachkov/ToDooo/internal/core/domain"
	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_request "github.com/MaximKachkov/ToDooo/internal/core/transport/http/request"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TasksCreated              int      `json:"tasks_created"`
	TasksCompleted            int      `json:"tasks_completed"`
	TaskCompletedRate         *float64 `json:"tasks_completion_rate"`
	TaskAverageComplteionTime *string  `json:"task_average_complteion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userId, fromTime, toTime, err := GetUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "not able to get query params in GetStatistics")
	}

	statDomain, err := h.statisticsService.GetStatistics(ctx, userId, fromTime, toTime)
	if err != nil {
		responseHandler.ErrorResponse(err, "error getting statistics")
		return
	}

	response := toDTOFromDomain(statDomain)

	responseHandler.JSONResponse(response, http.StatusOK)

}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.TaskAverageComplteionTime != nil {
		duration := statistics.TaskAverageComplteionTime.String()
		avgTime = &duration
	}

	return GetStatisticsResponse{
		TasksCreated:              statistics.TasksCreated,
		TasksCompleted:            statistics.TasksCompleted,
		TaskCompletedRate:         statistics.TaskCompletedRate,
		TaskAverageComplteionTime: avgTime,
	}
}

func GetUserIDFromToQueryParams(r *http.Request) (id *int, from *time.Time, to *time.Time, err error) {
	const (
		userIDQueryParamKey = "user_id"
		FromQueryParamKey   = "from"
		ToQueryParamKey     = "to"
	)
	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error in getting user id:%w ", err)
	}
	fromTime, err := core_http_request.GetTimeQueryParam(r, FromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error in getting from time:%w ", err)
	}
	toTime, err := core_http_request.GetTimeQueryParam(r, ToQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error in getting to time:%w ", err)
	}

	return userID, fromTime, toTime, nil
}
