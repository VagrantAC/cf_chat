package main

import (
	"cf_chat/route"
	"cf_chat/utils/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
  gin.SetMode(gin.ReleaseMode)
  server := gin.Default()

  server.Use(middleware.Cors())

  route.NewRoute(server)
  
  srv := &http.Server{
  	Addr:    ":8080",
  	Handler: server,
  }
  
  go func() {
  	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
  		log.Fatalf("listen: %s\n", err)
  	}
  }()

  quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
