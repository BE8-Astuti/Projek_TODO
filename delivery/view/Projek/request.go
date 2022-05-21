package projek


type InsertPro struct {
	UserID      uint   `json:"userid" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Contributor string `json:"contributor" validate:"required"`
}


type UpdatePro struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Contributor string `json:"contributor" validate:"required"`
}