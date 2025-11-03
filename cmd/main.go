package main

import (
	"log/slog"
	"os"

	server "github.com/cnt-777/internal/interfaces/http"
	"github.com/cnt-777/pkg/log"
)

func main() {
	lCfg := &log.Config{
		OutputType: log.Console,
		Level:      slog.LevelDebug,
	}

	handler, err := log.NewHandler(lCfg)
	if err != nil {
		slog.Error("failed create logger handler", slog.String("err", err.Error()))
		os.Exit(1)
	}
	l := slog.New(handler)

	l.Debug("logger inited")

	r := server.New()

	l.Info("server started")
	r.Run("0.0.0.0:8080")
}
