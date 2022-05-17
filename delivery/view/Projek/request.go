package projek

type InsertPro struct {
	UserID      uint   `json:"userid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Contributor string `json:"contributor"`
}

type UpdatePro struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Contributor string `json:"contributor"`
}
