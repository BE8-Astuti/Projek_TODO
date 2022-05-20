package task

type InsertTaskRequest struct {
	UserID   uint   `json:"userid" validate:"required"`
	ProjekID uint   `json:"projekid" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UpdateTaskRequest struct {
	Name    string `json:"name" validate:"required"`
	Status  string `json:"status"`
	Duedate string `json:"duedate"`
}
