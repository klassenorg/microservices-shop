package tests

import (
	v1 "cartsvc/internal/delivery/http/v1"
	"cartsvc/internal/repository"
	"cartsvc/internal/service"
	"cartsvc/pkg/db/redisdb"
	log "cartsvc/pkg/logger"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"os"
	"testing"
	"time"
)

var dbAddr string

func init() {
	dbAddr = os.Getenv("TEST_DB_URI")
}

type IntegrationTestSuite struct {
	suite.Suite

	db          *redis.Client
	repos       *repository.Repositories
	services    *service.Services
	httpHandler *v1.Handler
	//TODO: test grpcHandler
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	if client, err := redisdb.NewClient(dbAddr, "", 5*time.Second); err != nil {
		s.FailNow("failed to connect to db", err)
	} else {
		s.db = client
	}

	s.repos = repository.NewRepositories(s.db)
	s.services = service.NewServices(service.Deps{Repos: s.repos})

	l := zap.New(nil).Sugar()
	logger := &log.Logger{SugaredLogger: l}
	s.httpHandler = v1.NewHandler(s.services, logger)

	if err := s.populateDB(); err != nil {
		s.FailNow("failed to populate db", err)
	}
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.db.Close() //nolint:errcheck
}

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func (s *IntegrationTestSuite) populateDB() error {
	if err := s.db.HIncrBy(context.Background(), testGetCartUserID, testProductID, testBigCount).Err(); err != nil {
		return err
	}
	if err := s.db.HIncrBy(context.Background(), testRemoveFromCartUserID, testProductID, testBigCount).Err(); err != nil {
		return err
	}
	if err := s.db.HIncrBy(context.Background(), testCleanCartUserID, testProductID, testBigCount).Err(); err != nil {
		return err
	}

	return nil
}
