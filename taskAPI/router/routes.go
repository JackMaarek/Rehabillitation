package router

import (
	"fmt"
	"github.com/JackMaarek/Rehabillitation/taskAPI/controllers"
	"github.com/JackMaarek/Rehabillitation/taskAPI/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/tasks", controllers.GetAllTasks)
		v1.POST("/task", controllers.CreateTask)
	}

}

func Run(r *gin.Engine) {

	go func() {
		if err := r.Run(fmt.Sprintf(":%s", handlers.GetVariable("PORT"))); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// ----------------- CLOSE APP -----------------
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")
}
