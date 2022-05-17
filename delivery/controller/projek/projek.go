package projek

import (
	"net/http"
	middlewares "projek/todo/delivery/middleware"
	"projek/todo/delivery/view"
	projekV "projek/todo/delivery/view/projek"
	"projek/todo/entities"
	"projek/todo/repository/projek"
	"strconv"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ControllerProjek struct {
	Repo  projek.ProjekRepository
	Valid *validator.Validate
}

func NewControlPro(NewPro projek.ProjekRepository, validate *validator.Validate) *ControllerProjek {
	return &ControllerProjek{
		Repo:  NewPro,
		Valid: validate,
	}
}

// ADD NEW CART
func (p *ControllerProjek) CreateProjek() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Insert projekV.InsertPro
		if err := c.Bind(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := p.Valid.Struct(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		NewAdd := entities.Projek{
			UserID:      uint(UserID),
			Title:       Insert.Title,
			Description: Insert.Description,
			Contributor: Insert.Contributor,
		}
		result, errCreate := p.Repo.CreateProjek(NewAdd, uint(UserID))
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusCreated, projekV.SuccessInsert(result))
	}
}

func (p *ControllerProjek) GetAllProjek() echo.HandlerFunc {
	return func(c echo.Context) error {

		UserID := middlewares.ExtractTokenUserId(c)
		result, err := p.Repo.GetAllProjek(uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, projekV.StatusGetAllOk(result))
	}
}

func (p *ControllerProjek) GetProjekID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		idtask, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := p.Repo.GetProjekID(uint(idtask), uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, projekV.StatusGetIdOk(result))
	}
}

func (p *ControllerProjek) UpdateProjek() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		var update projekV.UpdatePro
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		UpdateTask := entities.Projek{
			Title:       update.Title,
			Description: update.Description,
			Contributor: update.Contributor,
		}

		result, errNotFound := p.Repo.UpdateProjek(uint(idcat), UpdateTask, uint(UserID))
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, projekV.StatusUpdate(result))
	}
}
func (p *ControllerProjek) DeleteProjek() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		taskid, err := strconv.Atoi(id)

		if err != nil {

			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		errDelete := p.Repo.DeleteProjek(uint(taskid), uint(UserID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
