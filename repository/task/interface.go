package task

import "projek/todo/entities"

type RepoTask interface {
	CreateTask(newAdd entities.Task, userid uint) (entities.Task, error)
	GetAllTask(Userid uint) ([]entities.Task, error)
	GetTaskID(id uint, userid uint) (entities.Task, error)
	UpdateTask(id uint, updatedAddress entities.Task, userid uint) (entities.Task, error)
	DeleteTask(id uint, Userid uint) error
	// SetDefaultProjekId(id uint, UserID uint) error
}
