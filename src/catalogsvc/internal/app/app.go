package app

import (
	"catalogsvc/internal/config"
	httpDelivery "catalogsvc/internal/delivery/http"
	"catalogsvc/internal/repository"
	"catalogsvc/internal/server"
	"catalogsvc/internal/service"
	"catalogsvc/pkg/db/mongodb"
	log "catalogsvc/pkg/logger"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	log.Init()
	logger := log.NewLogger()
	defer logger.Sync()
	logger.Info("logger initialized")

	cfg := config.Init()
	logger.Info("config initialized")

	mongoClient, err := mongodb.NewClient(cfg.MongoTimeout, cfg.MongoURI, cfg.MongoUser, cfg.MongoPassword)
	if err != nil {
		logger.Fatal(err)
	}

	db := mongoClient.Database(cfg.MongoDBName)
	logger.Infow("database initialized",
		"uri", cfg.MongoURI,
		"dbName", cfg.MongoDBName)

	repos := repository.NewRepositories(db)

	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	handlers := httpDelivery.NewHandler(services, logger)

	srv := server.NewServer(cfg, handlers.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Infow("server started",
		"port", cfg.Port)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
	}
}
