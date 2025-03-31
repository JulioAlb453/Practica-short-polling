package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LongPollingController(ctx *gin.Context) {
	lastKnownUpdate := ctx.Query("last_update")
	var lastKnownVal int
	if lastKnownUpdate != "" {
		fmt.Sscanf(lastKnownUpdate, "%d", &lastKnownVal)
	}

	updateChannel := make(chan int)
	go func() {
		for {
			if lastKnownVal != lastUpdate {
				updateChannel <- lastUpdate
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	select {
	case newUpdate := <-updateChannel:
		ctx.JSON(http.StatusOK, gin.H{
			"message":     "Se detectó un cambio en la información.",
			"Hora de la ultima actualizacion": newUpdate,
		})
	case <-time.After(30 * time.Second):
		ctx.JSON(http.StatusNotModified, gin.H{
			"message": "No se realizaron cambios en el intervalo de espera.",
		})
	}
}
