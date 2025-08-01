package main

import (
	"github.com/JackMaarek/Rehabillitation/taskAPI/db"
	"github.com/JackMaarek/Rehabillitation/taskAPI/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// init db connection
	if err := db.Init(); err != nil {
		log.Fatal("Could not connect to db : ", err)
	}
	db.Migrate()

	// init router
	r := gin.Default()
	router.Initialize(r)
	router.Run(r)
	log.Info("✅ API running & Listening")
}
