package task

import "time"

type InsertTaskRequest struct {
	UserID   uint      `json:"user_id" validate:"required"`
	ProjekID uint      `json:"projek_id" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Duedate  time.Time `json:"duedate"`
}

type UpdateTaskRequest struct {
	Name    string    `json:"name" validate:"required"`
	Status  string    `json:"status"`
	Duedate time.Time `json:"duedate"`
}
