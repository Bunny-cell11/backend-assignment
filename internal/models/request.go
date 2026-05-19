package models

type RequestPayload struct {
	UserID  string      `json:"user_id"`
	Payload interface{} `json:"payload"`
}

type UserStats struct {
	Accepted int `json:"accepted_requests_current_window"`
	Rejected int `json:"rejected_requests_total"`
}
