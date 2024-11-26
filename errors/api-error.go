package errors

type ApiError struct {
	Error string `json:"error"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
