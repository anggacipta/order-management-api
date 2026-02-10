package dto

type LogRequest struct {
	Action    string `json:"action" binding:"required"`
	UserID    uint   `json:"user_id" binding:"required"`
	Timestamp int64  `json:"timestamp" binding:"required"`
}
