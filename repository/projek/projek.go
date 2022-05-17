package projek

import (
	"errors"
	"projek/todo/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ProjekDB struct {
	Db *gorm.DB
}

// Get Access DB
func NewDB(db *gorm.DB) *ProjekDB {
	return &ProjekDB{
		Db: db,
	}
}

func (p *ProjekDB) CreateProjek(newAdd entities.Projek, userid uint) (entities.Projek, error) {
	if err := p.Db.Where("userid= ?", userid).Create(&newAdd).Error; err != nil {
		log.Warn(err)
		return newAdd, err
	}
	return newAdd, nil
}
func (p *ProjekDB) GetAllProjek(userid uint) ([]entities.Projek, error) {
	arrprojek := []entities.Projek{}

	if err := p.Db.Where("user_id= ?", userid).Find(&arrprojek).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select data")
	}

	if len(arrprojek) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrprojek, nil
}

func (p *ProjekDB) GetProjekID(id uint, userid uint) (entities.Projek, error) {
	var projek entities.Projek
	if err := p.Db.Where("id= ? AND user_id= ?", id, userid).Find(&projek).Error; err != nil {
		log.Warn("Error Get By ID", err)
		return projek, err
	}
	return projek, nil
}

func (p *ProjekDB) UpdateProjek(id uint, UpdatePro entities.Projek, UserID uint) (entities.Projek, error) {
	var projek entities.Projek

	if err := p.Db.Where("id =? AND user_id =?", id, UserID).First(&projek).Updates(&UpdatePro).Find(&projek).Error; err != nil {
		log.Warn("Update Error", err)
		return projek, err
	}

	return projek, nil
}

func (p *ProjekDB) DeleteProjek(id uint, UserID uint) error {

	var delete entities.Projek
	if err := p.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}
