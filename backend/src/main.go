package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/y-watagashi/MoneyCraft/utils"
	"github.com/y-watagashi/MoneyCraft/models"
	"github.com/y-watagashi/MoneyCraft/controllers"
)

func main() {
	r := gin.Default()
	db := utils.ConnectToDB()
	err := db.AutoMigrate(&models.Family{}, &models.Transaction{}, &models.User{}, &models.History{})
	if err != nil {
        fmt.Println("failed to migrate:", err)
    }

	userHandler := controllers.UserHandler{Db: db}
	historyHandler := controllers.HistoryHandler{Db: db}

	// User関連のエンドポイント
	r.POST("/user", userHandler.AddUser)
	r.GET("/user/:id", userHandler.GetUser)

	// History関連のエンドポイント
	r.POST("/history/:id", historyHandler.AddHistory)
	
	r.Run("0.0.0.0:8000") // listen and serve on
}