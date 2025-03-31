package server

import (
	"net/http"
	"time"
)

func Run() {
	r := gin.Default()


	srv := &http.Server{
		Addr:         ":4000",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic("Error: Server Main hasn't begun")
	}
}
