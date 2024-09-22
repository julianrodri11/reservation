package dto

import "time"

type ReservationDTO struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	FieldID   uint      `json:"field_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
