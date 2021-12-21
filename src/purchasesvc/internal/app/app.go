package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"purchasesvc/internal/config"
	"purchasesvc/internal/repository"
	"purchasesvc/internal/server"
	"purchasesvc/internal/service"
	"purchasesvc/pkg/cartclient"
	"purchasesvc/pkg/logger"
	"purchasesvc/pkg/paymentclient"
	"purchasesvc/pkg/pricingclient"
	"runtime"
	"syscall"
	"time"

	grpcDelivery "purchasesvc/internal/delivery/grpc"
	httpDelivery "purchasesvc/internal/delivery/http"
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

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresUser,
			cfg.PostgresDBName,
			cfg.PostgresPassword,
			cfg.PostgresSSLMode),
	)
	if err != nil {
		log.Fatalw("error initializing db",
			"error", err)
	}

	repos := repository.NewRepositories(db)

	cartClient, err := cartclient.NewCartClient(cfg.CartSvcGRPCAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(cartClient *cartclient.Client) {
		err := cartClient.Close()
		if err != nil {
			log.Errorw("error closing catalog client",
				"error", err)
		}
	}(cartClient)

	pricingClient, err := pricingclient.NewPricingClient(cfg.PricingSvcGRPCAddr)
	if err != nil {
		log.Fatalw("error initializing pricing client",
			"error", err)
	}
	defer func(pricingClient *pricingclient.Client) {
		err := pricingClient.Close()
		if err != nil {
			log.Errorw("error closing pricing client",
				"error", err)
		}
	}(pricingClient)

	paymentProvider := paymentclient.NewClient()

	services := service.NewServices(service.Deps{
		Repos:           repos,
		CartClient:      cartClient,
		PricingClient:   pricingClient,
		PaymentProvider: paymentProvider,
	})

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
