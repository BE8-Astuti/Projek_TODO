package task

import "github.com/labstack/echo/v4"

type TaskControl interface {
	CreateTask() echo.HandlerFunc
	UpdateTask() echo.HandlerFunc
	GetAllTask() echo.HandlerFunc
	GetTaskID() echo.HandlerFunc
	DeleteTask() echo.HandlerFunc
}
