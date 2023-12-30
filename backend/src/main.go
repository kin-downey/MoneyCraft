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
	db.AutoMigrate(&models.User{})

	userHandler := controllers.UserHandler{Db: db}
	r.POST("/user", userHandler.AddUser)
	r.GET("/user/:id", userHandler.GetUser)
	
	r.Run("0.0.0.0:8000") // listen and serve on
}