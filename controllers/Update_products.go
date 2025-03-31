package controllers

import (
	"Practica/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateProductController(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var updateData models.Productos
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Información no válida"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			product.Nombre = updateData.Nombre
			product.Cant = updateData.Cant
			product.CodigoBarras = updateData.CodigoBarras
			product.UpdatedAt = time.Now()

			lastUpdate = int(time.Now().UnixNano())


			products[i] = product

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Producto actualizado correctamente.",
				"product": product,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
}

