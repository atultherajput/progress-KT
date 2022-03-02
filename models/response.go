package models

type JsonResponse struct {
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
