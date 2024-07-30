package main

import (
	"gin-template/internal/server/app"
	"gin-template/internal/server/config"
	"gin-template/lib/prettylog"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := prettylog.SetupLogger(cfg.Env)

	log.Info("initializing...",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.Server.Port),
		slog.Any("config", cfg),
	)

	application := app.NewApp(log, cfg)

	log.Info("starting ...")
	go application.MustRun()

	// --------------------------- Register stop signal ---------------------------
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// --------------------------- Wait for stop signal ---------------------------
	sign := <-stop

	log.Info("shutting down ...",
		slog.String("signal", sign.String()),
	)
	err := application.Stop()
	if err != nil {
		log.Error("failed to stop the application", slog.Any("error", err))
	}

	log.Info("application fully stopped")

}
