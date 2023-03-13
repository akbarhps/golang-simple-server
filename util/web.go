package util

type WebResponse struct {
	StatusCode int         `json:"status_code,omitempty"`
	Status     string      `json:"status,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
