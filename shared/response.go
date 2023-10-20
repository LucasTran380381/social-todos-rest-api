package shared

import "time"

type Response struct {
	Data      interface{} `json:"data"`
	Paging    *Paging     `json:"paging,omitempty"`
	Filter    interface{} `json:"filter,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

type ErrorResponse struct {
	Error     string    `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}
