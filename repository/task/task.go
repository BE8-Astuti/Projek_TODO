package task

import (
	"projek/todo/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type TaskDB struct {
	Db *gorm.DB
}

// Get Access DB
func NewDB(db *gorm.DB) *TaskDB {
	return &TaskDB{
		Db: db,
	}
}

func (t *TaskDB) CreateTask(newAdd entities.Task, userid uint) (entities.Task, error) {
	if err := t.Db.Where("userid= ?", userid).Create(&newAdd).Error; err != nil {
		log.Warn(err)
		return newAdd, err
	}
	return newAdd, nil
}
func (t *TaskDB) GetAllTask(userid uint) ([]entities.Task, error) {
	var task []entities.Task
	if err := t.Db.Where("user_id= ?", userid).Find(&task).Error; err != nil {
		log.Warn("Error Get Data", err)
		return task, err
	}
	return task, nil
}

func (t *TaskDB) GetTaskID(id uint, userid uint) (entities.Task, error) {
	var task entities.Task
	if err := t.Db.Where("id= ? AND user_id= ?", id, userid).Find(&task).Error; err != nil {
		log.Warn("Error Get By ID", err)
		return task, err
	}
	return task, nil
}

func (t *TaskDB) UpdateTask(id uint, UpdateTask entities.Task, UserID uint) (entities.Task, error) {
	var task entities.Task

	if err := t.Db.Where("id =? AND user_id =?", id, UserID).First(&task).Updates(&UpdateTask).Find(&task).Error; err != nil {
		log.Warn("Update Error", err)
		return task, err
	}

	return task, nil
}

func (t *TaskDB) DeleteTask(id uint, UserID uint) error {

	var delete entities.Task
	if err := t.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}
