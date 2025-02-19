package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tarun222999/students-api/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//logger
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})
	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("server started at %s", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failer to start server due to %s", err)
	}

}
