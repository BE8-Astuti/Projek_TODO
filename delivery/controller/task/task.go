package task

import (
	"net/http"
	middlewares "projek/todo/delivery/middleware"
	"projek/todo/delivery/view"
	taskV "projek/todo/delivery/view/task"
	"projek/todo/entities"
	"projek/todo/repository/task"
	"strconv"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ControlTask struct {
	Repo  task.RepoTask
	Valid *validator.Validate
}

func NewControlTask(NewTask task.RepoTask, validate *validator.Validate) *ControlTask {
	return &ControlTask{
		Repo:  NewTask,
		Valid: validate,
	}
}

// ADD NEW CART
func (t *ControlTask) CreateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Insert taskV.InsertTaskRequest
		if err := c.Bind(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := t.Valid.Struct(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		NewAdd := entities.Task{
			UserID:   uint(UserID),
			ProjekID: Insert.ProjekID,
			Name:     Insert.Name,
		}
		result, errCreate := t.Repo.CreateTask(NewAdd, uint(UserID))
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusCreated, taskV.SuccessInsert(result))
	}
}

func (t *ControlTask) GetAllTask() echo.HandlerFunc {
	return func(c echo.Context) error {

		UserID := middlewares.ExtractTokenUserId(c)
		result, err := t.Repo.GetAllTask(uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, taskV.StatusGetAllOk(result))
	}
}

func (t *ControlTask) GetTaskID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		idtask, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := t.Repo.GetTaskID(uint(idtask), uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, taskV.StatusGetIdOk(result))
	}
}

func (t *ControlTask) UpdateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		var update taskV.UpdateTaskRequest
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		UpdateTask := entities.Task{
			Name:   update.Name,
			Status: update.Status,
		}

		result, errNotFound := t.Repo.UpdateTask(uint(idcat), UpdateTask, uint(UserID))
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, taskV.StatusUpdate(result))
	}
}
func (t *ControlTask) DeleteTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		taskid, err := strconv.Atoi(id)

		if err != nil {

			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		errDelete := t.Repo.DeleteTask(uint(taskid), uint(UserID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
