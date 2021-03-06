package controllers

import (
	"gingorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name  string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
}

type UpdateProductInput struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

// GET /products
// Find all books
func FindProducts(c *gin.Context) {
	var products []model.Product
	model.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GET /products/:id
// Find a product
func FindProduct(c *gin.Context) {
	// Get model if exist
	var product model.Product
	if err := model.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// POST /products
// Create new product
func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create product
	product := model.Product{Name: input.Name, Price: input.Price}
	model.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// PATCH /products/:id
// Update a product
func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product model.Product
	if err := model.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.DB.Model(&product).Updates(product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DELETE /products/:id
// Delete a products
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var product model.Product
	if err := model.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	model.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
