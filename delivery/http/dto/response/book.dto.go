package response

type ResponseErrorDTO struct {
	StatusCode int `json:"status_code"`
	Error      any `json:"error"`
}

type ResponseSuccessDTO struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
