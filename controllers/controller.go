package controllers

import (
	"go-rest/config"
	"go-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	DB := config.FetchDB()

	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		if err := DB.Create(&order).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"order": order})
		}
	}
}

func GetOrders(c *gin.Context) {
	DB := config.FetchDB()

	orders := []models.Order{}
	if err := DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"orders": orders})
	}
}
