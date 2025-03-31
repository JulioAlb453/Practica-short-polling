package controllers

import (
	"Practica/models"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
)

var products []models.Productos
var lastUpdate int

func CreateProductsHandler(ctx context.Context){
	var product models.Productos

	if err := ctx.ShouldBindJSON(&product); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Informacion no valida"})
        return
	}

	product.ID = len(products) + 1
	products = append(products, product)
	
	lastUpdate = product.ID

	ctx.JSON(http.StatusCreated, gin.H{
		"message":  "Producto creado exitosamente",
		"producto": product, 
	})
}