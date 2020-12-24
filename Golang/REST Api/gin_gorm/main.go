package main

import (
	//"net/http"
	"gingorm/controllers"
	"gingorm/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	model.ConnectDatabase()

	// Routes
	r.GET("/products", controllers.FindProducts)
	r.GET("/products/:id", controllers.FindProduct)
	r.POST("/products", controllers.CreateProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	// Run the server
	r.Run()
}
