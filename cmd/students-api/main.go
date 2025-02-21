package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Tarun222999/students-api/internal/config"
	student "github.com/Tarun222999/students-api/internal/http/handlers/students"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//logger
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())
	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server starting ", slog.String("address", cfg.HTTPServer.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server due to %s", err)
		}
	}()

	<-done

	slog.Info("shutting down the server ")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to stop server due to ", slog.String("error", err.Error()))
	}
}
