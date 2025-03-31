package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductsController (ctx *gin.Context){
	ctx.JSON(http.StatusOK, products)
}