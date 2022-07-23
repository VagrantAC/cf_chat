package main

import (
	"cf_chat/framework"
	"cf_chat/framework/middleware"
	"cf_chat/route"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Cost(), middleware.Recovery())
	route.RegisterRouter(core)
	server := &http.Server{
    Handler: core, 
    Addr:    ":8000",
  }
  go func() {
 		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	} else {
		log.Println("Server Shutdown Success")
	}
}

