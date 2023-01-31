package controller

import (
	"errors"
	database "github.com/Lefree111/gorm-resta-api/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateData(c *gin.Context) {
	var book *database.Data
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
}

func ReadData(c *gin.Context) {
	var book database.Data
	id := c.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
}

func ReadDatas(c *gin.Context) {
	var books []database.Data
	res := database.DB.Find(&books)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("authors not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
	return
}

func UpdateData(c *gin.Context) {
	var data database.Data
	id := c.Param("id")
	err := c.ShouldBind(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateData database.Data
	res := database.DB.Model(&updateData).Where("id = ?", id).Updates(data)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "book not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": data,
	})
	return
}

func DeleteData(c *gin.Context) {
	var data database.Data
	id := c.Param("id")
	res := database.DB.Find(&data, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	database.DB.Delete(&data)
	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
	return
}
