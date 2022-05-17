package routes

import (
	"projek/todo/delivery/controller/projek"
	"projek/todo/delivery/controller/task"
	"projek/todo/delivery/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo, u user.ControllerUser, t task.TaskControl, p projek.ProjekControl) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// Login
	e.POST("/login", u.Login())
	// ROUTES USER
	user := e.Group("/user")
	user.POST("", u.InsertUser()) // Register
	// user.GET("", u.GetAllUser, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.GET("/:id", u.GetUserbyID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	user.PUT("/:id", u.UpdateUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	user.DELETE("/:id", u.DeleteUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))

	task := e.Group("/task")
	task.POST("", t.CreateTask(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	task.GET("", t.GetAllTask(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	task.GET("/:id", t.GetTaskID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	task.PUT("/:id", t.UpdateTask(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	task.DELETE("/:id", t.DeleteTask(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))

	projek := e.Group("/projek")
	projek.POST("", p.CreateProjek(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	projek.GET("", p.GetAllProjek(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	projek.GET("/:id", p.GetProjekID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	projek.PUT("/:id", p.UpdateProjek(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
	projek.DELETE("/:id", p.DeleteProjek(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TODO")}))
}
