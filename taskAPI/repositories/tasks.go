package repositories

import (
	"github.com/JackMaarek/Rehabillitation/taskAPI/db"
	"github.com/JackMaarek/Rehabillitation/taskAPI/models"
)

func FindAllTasks(tasks *[]models.Task) error {
	return db.Db.Debug().Find(tasks).Error
}

func FindTaskByName(task *models.Task) error {
	return db.Db.Debug().Where("name = ?", task.Name).First(&task).Error
}

// CreateTask persist server in database
func CreateTask(task *models.Task) error {
	return db.Db.Debug().Create(task).Error
}