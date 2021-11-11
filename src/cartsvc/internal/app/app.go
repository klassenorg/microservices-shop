package app

import (
	"cartsvc/internal/config"
	grpcDelivery "cartsvc/internal/delivery/grpc"
	httpDelivery "cartsvc/internal/delivery/http"
	"cartsvc/internal/repository"
	"cartsvc/internal/server"
	"cartsvc/internal/service"
	"cartsvc/pkg/db/redisdb"
	log "cartsvc/pkg/logger"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func Run() {
	log.Init()
	logger := log.NewLogger()
	defer func(logger *log.Logger) {
		_ = logger.Sync()
	}(logger)
	logger.Info("logger initialized")

	cfg := config.Init()
	logger.Info("config initialized")

	//pprof server
	if cfg.Debug {
		go func() {
			runtime.SetBlockProfileRate(1)
			if err := http.ListenAndServe(":"+cfg.DebugPprofPort, nil); !errors.Is(err, http.ErrServerClosed) {
				logger.Errorf("error ocurried while running hprof server: %s\n", err.Error())
			}
		}()
		logger.Infow("hprof server started",
			"port", cfg.DebugPprofPort)
	}

	redisClient, err := redisdb.NewClient(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisTimeout)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infow("database initialized",
		"addr", cfg.RedisAddr)

	repos := repository.NewRepositories(redisClient)

	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handlers := httpDelivery.NewHandler(services, logger)

	grpcHandlers := grpcDelivery.NewHandler(services, logger)

	grpcSrv, err := server.NewGRPCServer(cfg, grpcHandlers)
	if err != nil {
		logger.Errorf("error starting GRPC server: %s\n", err.Error())
	}

	go func() {
		if err := grpcSrv.Run(); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			logger.Errorf("error occurried while running grpc server: %s\n", err.Error())
		}
	}()

	logger.Infow("grpc server started",
		"port", cfg.GRPCPort)

	srv := server.NewServer(cfg, handlers.Init())

	go func() {
		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurried while running http server: %s\n", err.Error())
		}
	}()

	logger.Infow("http server started",
		"port", cfg.HTTPPort)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	logger.Info("service started successfully")

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	grpcSrv.Stop()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := redisClient.Close(); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("service stopped successfully")

}
