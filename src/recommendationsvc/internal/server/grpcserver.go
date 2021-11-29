package server

import (
	"google.golang.org/grpc"
	"net"
	pb "recommendationsvc/gen/proto"
	"recommendationsvc/internal/config"
)

type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg *config.Config, srv pb.RecommendationServiceServer) (*GRPCServer, error) {
	listener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		return nil, err
	}

	registrar := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(registrar, srv)

	return &GRPCServer{listener: listener, server: registrar}, nil
}

func (s *GRPCServer) Run() error {
	return s.server.Serve(s.listener)
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}
