package core_http_response

type ErrorResponse struct {
	Message string `json:"message"    example:"hello"`
	Error   string `json:"error"      example:"short human-readable message"`
}
