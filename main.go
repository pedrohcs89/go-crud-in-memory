package main

import (
	"crud-go/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("faild ro execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("all system offline")
}

func run() error {
	db := make(map[string]api.User)
	handdler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handdler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
