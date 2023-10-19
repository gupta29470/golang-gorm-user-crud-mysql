package models

type ErrorResponse struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}
