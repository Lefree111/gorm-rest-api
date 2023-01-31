package main

import (
	"fmt"

	database "github.com/Lefree111/gorm-resta-api/api/entity"
	"github.com/Lefree111/gorm-resta-api/rest-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/books/:id", controller.ReadData)
	r.GET("/books", controller.ReadDatas)
	r.POST("/books", controller.CreateData)
	r.PUT("/books/:id", controller.UpdateData)
	r.DELETE("/books/:id", controller.DeleteData)
	r.Run(":5000")
}
