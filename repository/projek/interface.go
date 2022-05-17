package projek

import "projek/todo/entities"

type ProjekRepository interface {
	CreateProjek(newAdd entities.Projek, userid uint) (entities.Projek, error)
	GetAllProjek(userid uint) ([]entities.Projek, error)
	GetProjekID(id uint, userid uint) (entities.Projek, error)
	UpdateProjek(id uint, UpdatePro entities.Projek, userid uint) (entities.Projek, error)
	DeleteProjek(id uint, userid uint) error
}
