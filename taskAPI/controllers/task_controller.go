package controllers

import (
	"fmt"
	"github.com/JackMaarek/Rehabillitation/taskAPI/models"
	"github.com/JackMaarek/Rehabillitation/taskAPI/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, "request should contain valid properties")
		return
	}

	foundTask := models.Task{
		Name: task.Name,
	}

	if err := repositories.FindTaskByName(&foundTask); err == nil {
		c.JSON(http.StatusOK, fmt.Sprintf("Task already exists: %s", foundTask.Name))
		return
	}

	if err := repositories.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Could not create birthday for user id %s", task.Name))
		return
	}

	c.JSON(http.StatusOK, "User's Task added successfully")
}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	if err := repositories.FindAllTasks(&tasks); err != nil {
		c.JSON(http.StatusInternalServerError, "error while fetching tasks")
		return
	}

	c.JSON(http.StatusOK, tasks)
}