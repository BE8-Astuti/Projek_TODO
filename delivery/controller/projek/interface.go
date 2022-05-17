package projek

import "github.com/labstack/echo/v4"

type ProjekControl interface {
	CreateProjek() echo.HandlerFunc
	UpdateProjek() echo.HandlerFunc
	GetAllProjek() echo.HandlerFunc
	GetProjekID() echo.HandlerFunc
	DeleteProjek() echo.HandlerFunc
}
