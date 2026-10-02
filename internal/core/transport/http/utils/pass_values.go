package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func GetInPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("no key='%s' in path value :%w", key, core_errors.ErrInvalidArgument)
	}
	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf("invalid pathvalue='%s' in path value :%w", pathValue, core_errors.ErrInvalidArgument)
	}
	return val, nil
}
