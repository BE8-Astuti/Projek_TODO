package task

import (
	"net/http"
)

type RespondTask struct {
	UserID   uint   `json:"userid"`
	ProjekID uint   `json:"projekid"`
	Name     string `json:"name"`
	Duedate  string `json:"duedate"`
	Status   string `json:"status"`
}

func StatusGetAllOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All data",
		"status":  true,
		"data":    data,
	}
}

func SuccessInsert(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Task",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Bad Request Access",
		"status":  false,
		"data":    nil,
	}
}

func StatusUpdate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
func StatusGetIdOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data ID",
		"status":  true,
		"data":    data,
	}
}
