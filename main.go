package main

import (
	"cf_chat/framework"
	"cf_chat/route"
	"cf_chat/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Cost(), middleware.Recovery())
	route.RegisterRouter(core)
	server := &http.Server{
    Handler: core, 
    Addr:    ":8000",
  }
  server.ListenAndServe()
}

