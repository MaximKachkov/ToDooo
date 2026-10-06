package core_http_response

import "net/http"

var (
	StatusCodeUnitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUnitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(statuscode int) {
	rw.ResponseWriter.WriteHeader(statuscode)
	rw.statusCode = statuscode
}

func (rw *ResponseWriter) GetStatusCode() int {
	if rw.statusCode == StatusCodeUnitialized {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.statusCode
}
