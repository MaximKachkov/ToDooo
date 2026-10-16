package web_transport_http

import (
	"net/http"

	core_logger "github.com/MaximKachkov/ToDooo/internal/core/logger"
	core_http_response "github.com/MaximKachkov/ToDooo/internal/core/transport/http/response"
)

func (h *WebHTTPHandler) GetMainPage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	html, err := h.webService.GetMainPage()
	if err != nil {
		responseHandler.ErrorResponse(err, "wasnt able to get main page")
		return
	}

	responseHandler.HTMLResponse(html)

}
