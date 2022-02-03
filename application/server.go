package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT += "8080"
	}
	router := SetupRouter()

	server := &http.Server{
		Addr: PORT,
		Handler: router,
	}
	wait := make(chan os.Signal)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Printf("Oops!!!, An error occurred: %v\n", err.Error())
        }
	}()
	log.Printf("Server started at localhost%v\n", PORT)
	signal.Notify(wait, os.Interrupt)
	<-wait
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	log.Println("Server Shutting down...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error while shutting down server: %v", err)
	}
	time.Sleep(time.Second)
	log.Println("Server shutdown successfully")
}


