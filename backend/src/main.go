package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/y-watagashi/MoneyCraft/utils"
	"github.com/y-watagashi/MoneyCraft/controllers"
)

func main() {
	db := utils.ConnectToDB()
	fmt.Println(db)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.GET("/test", controllers.Test)
	r.Run("0.0.0.0:8000") // listen and serve on
}