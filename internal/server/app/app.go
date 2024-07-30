package app

import (
	"errors"
	"fmt"
	"gin-template/internal/server/config"
	handlers "gin-template/internal/server/handlers"
	"gin-template/internal/server/routes"
	"gin-template/internal/server/services/auth"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type App struct {
	Log *slog.Logger
	Srv *http.Server
}

func NewApp(log *slog.Logger, cfg *config.Config) *App {

	service := auth.NewService(log)

	authHandlerTmp := handlers.NewAuthHandler(log, service)
	h := &routes.Handlers{
		AuthHandler: authHandlerTmp,
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	routes.SetupRoutes(router, h)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", strconv.Itoa(cfg.Server.Port)),
		Handler: router,
	}

	return &App{
		Log: log,
		Srv: srv,
	}
}

func (a *App) run() error {
	if err := a.Srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
		return err
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() error {
	a.Log.Info("stopping")
	err := a.Srv.Shutdown(nil)
	if err != nil {
		return err
	}
	return nil
}
