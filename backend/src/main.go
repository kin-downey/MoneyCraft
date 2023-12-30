package main

import (
	"github.com/gin-gonic/gin"
	"github.com/y-watagashi/MoneyCraft/utils"
	"github.com/y-watagashi/MoneyCraft/models"
	"github.com/y-watagashi/MoneyCraft/controllers"
)

func main() {
	r := gin.Default()
	db := utils.ConnectToDB()
	err := db.AutoMigrate(&models.User{}, &models.History{})
	if err != nil {
        log.Fatal("failed to migrate:", err)
    }

	userHandler := controllers.UserHandler{Db: db}
	r.POST("/user", userHandler.AddUser)
	r.GET("/user/:id", userHandler.GetUser)
	
	r.Run("0.0.0.0:8000") // listen and serve on
}