package model

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
