package controllers

import (
	"Practica/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortPollingController(ctx *gin.Context) {
	var history []models.Productos

	for _, product := range products {
		if !product.UpdatedAt.IsZero() {
			history = append(history, product)
		}
	}

	if len(history) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No se han registrado actualizaciones.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Se han detectado las siguientes actualizaciones:",
		"updates": history,
	})
}
