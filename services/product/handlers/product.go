package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leedrum/supa-shop/services/product/db"
	"github.com/leedrum/supa-shop/services/product/models"
)

func Create(c *gin.Context) {
	product := models.Product{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product is not valid"})
		return
	}

	if db.DB.Create(&product).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not create the product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func Delete(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id can not be empty"})
		return
	}

	if db.DB.Delete(&models.Product{}, id).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can not delete the product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
