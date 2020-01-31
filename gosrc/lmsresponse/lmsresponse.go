package lmsresponse

import (
	"encoding/json"
)

// constants ...
const (
	SUCCESS = "success"
	ERROR   = "error"
)

// Response is a generic server response
type Response struct {
	Status  string      `json:"status" binding:"required"`
	Message string      `json:"message" binding:"required"`
	Data    interface{} `json:"data,omitempty"`
}

// GetResponseBytes ...
func GetResponseBytes(status string, message string, data interface{}) []byte {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	res, err := json.Marshal(response)
	if err != nil {
		res, _ = json.Marshal(Response{
			Status:  "error",
			Message: "Unable to marshal server response",
		})
		return res
	}
	return res
}
