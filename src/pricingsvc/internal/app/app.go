package app

import (
	"pricingsvc/internal/config"
	"pricingsvc/pkg/cartclient"
	"pricingsvc/pkg/catalogclient"

	//grpcDelivery "pricingsvc/internal/delivery/grpc"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	httpDelivery "pricingsvc/internal/delivery/http"
	"pricingsvc/internal/server"
	"pricingsvc/internal/service"
	log "pricingsvc/pkg/logger"
	//"google.golang.org/grpc"
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

	cartClient, err := cartclient.NewCartClient(cfg.CartSvcGRPCAddr)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(cartClient *cartclient.CartClient) {
		err := cartClient.Close()
		if err != nil {
			logger.Errorw("error closing cart client",
				"error", err)
		}
	}(cartClient)

	catalogClient, err := catalogclient.NewCartClient(cfg.CatalogSvcGRPCAddr)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(catalogClient *catalogclient.CatalogClient) {
		err := catalogClient.Close()
		if err != nil {
			logger.Errorw("error closing catalog client",
				"error", err)
		}
	}(catalogClient)

	services := service.NewServices(service.Deps{
		CartClient:    cartClient,
		CatalogClient: catalogClient,
	})

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handlers := httpDelivery.NewHandler(services, logger)

	//grpcHandlers := grpcDelivery.NewHandler(services, logger)

	//grpcSrv, err := server.NewGRPCServer(cfg, grpcHandlers)
	//if err != nil {
	//	logger.Errorf("error starting GRPC server: %s\n", err.Error())
	//}
	//
	//go func() {
	//	if err := grpcSrv.Run(); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
	//		logger.Errorf("error occurried while running grpc server: %s\n", err.Error())
	//	}
	//}()
	//
	//logger.Infow("grpc server started",
	//	"port", cfg.GRPCPort)

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

	//grpcSrv.Stop()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	logger.Info("service stopped successfully")

}
