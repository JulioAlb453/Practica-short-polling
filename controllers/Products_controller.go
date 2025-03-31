package controllers

import (
	"Practica/models"

	"github.com/gin-gonic/gin"
)

var products []models.Productos
var lastUpdate int

func CreateProductController(ctx *gin.Context) {
	var product models.Productos
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	product.ID = len(products) + 1
	products = append(products, product)

	

	ctx.JSON(200, gin.H{
		"message": "Producto creado exitosamente",
		"product": product,
	})
}
