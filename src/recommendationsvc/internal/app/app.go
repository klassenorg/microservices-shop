package app

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"recommendationsvc/internal/config"
	grpcDelivery "recommendationsvc/internal/delivery/grpc"
	httpDelivery "recommendationsvc/internal/delivery/http"
	"recommendationsvc/internal/server"
	"recommendationsvc/internal/service"
	"recommendationsvc/pkg/catalogclient"
	"recommendationsvc/pkg/logger"
	"runtime"
	"syscall"
	"time"
)

func Run() {
	logger.Init()
	log := logger.NewLogger()
	defer func(logger *logger.Logger) {
		_ = logger.Sync()
	}(log)
	log.Info("logger initialized")

	cfg := config.Init()
	log.Info("config initialized")

	if cfg.Debug {
		go func() {
			runtime.SetBlockProfileRate(1)
			if err := http.ListenAndServe(":"+cfg.DebugPprofPort, nil); !errors.Is(err, http.ErrServerClosed) {
				log.Errorf("error ocurried while running hprof server: %s\n", err.Error())
			}
		}()
		log.Infow("hprof server started",
			"port", cfg.DebugPprofPort)
	}

	catalogClient, err := catalogclient.NewCartClient(cfg.CatalogSvcGRPCAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(catalogClient *catalogclient.Client) {
		err := catalogClient.Close()
		if err != nil {
			log.Errorw("error closing catalog client",
				"error", err)
		}
	}(catalogClient)

	services := service.NewServices(service.Deps{CatalogClient: catalogClient})

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handlers := httpDelivery.NewHandler(services, log)

	grpcHandlers := grpcDelivery.NewHandler(services, log)

	grpcSrv, err := server.NewGRPCServer(cfg, grpcHandlers)
	if err != nil {
		log.Errorf("error starting GRPC server: %s\n", err.Error())
	}

	go func() {
		if err := grpcSrv.Run(); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Errorf("error occurried while running grpc server: %s\n", err.Error())
		}
	}()

	log.Infow("grpc server started",
		"port", cfg.GRPCPort)

	srv := server.NewServer(cfg, handlers.Init())

	go func() {
		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurried while running http server: %s\n", err.Error())
		}
	}()

	log.Infow("http server started",
		"port", cfg.HTTPPort)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	log.Info("service started successfully")

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	grpcSrv.Stop()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}

	log.Info("service stopped successfully")

}
