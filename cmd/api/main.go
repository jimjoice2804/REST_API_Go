package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jimjoice2804/REST_API_Go/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	//setup router
	route := http.NewServeMux()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go"))
	})
	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: route,
	}

	fmt.Println("welcome to api", cfg)
	slog.Info("server started", "address", cfg.Address)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown", slog.String("error", err.Error()))
	}
	slog.Info("shutdown successfully")
}
