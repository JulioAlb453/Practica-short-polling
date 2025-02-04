package server

import (
	"Practica/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var products []models.Productos
var lastUpdate int

func createProductsHandler(ctx *gin.Context) {

	var product models.Productos

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Informacion no valida"})
		return
	}
	product.ID = len(products) + 1
	products = append(products, product) 

	lastUpdate = product.ID 

	ctx.JSON(http.StatusCreated, product)

}

func getProductsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, products)
}

func getLastUpdateHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"last_update": lastUpdate})
}

func Run(){
	r := gin.Default()

	r.POST("/products", createProductsHandler)
	r.GET("/products", getProductsHandler)

	r.GET("/LastChange", getLastUpdateHandler)
	
	srv := &http.Server{
		Addr:         ":4000",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Error: Server Main hasn't begin")
	}
}
