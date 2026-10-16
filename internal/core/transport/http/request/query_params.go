package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}
	integer, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("param= %s , by key= %s , not a valid integer:%v ,%w", param, key, err, core_errors.ErrInvalidArgument)
	}
	return &integer, nil
}

func GetTimeQueryParam(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}
	parsedTime, err := time.Parse(time.RFC3339, param)
	if err != nil {
		parsedTime, err = time.Parse("2006-01-02", param)
	}
	if err != nil {
		return nil, fmt.Errorf("error parsing a time form param:%s , %v , %w", param, err, core_errors.ErrInvalidArgument)
	}
	return &parsedTime, nil

}
