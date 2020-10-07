package grpc

import (
	"log"
	"net"

	"github.com/sueken5/go-integration-test-example/pkg/apis/sample"
	"github.com/sueken5/go-integration-test-example/pkg/application"
	"github.com/sueken5/go-integration-test-example/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	lis net.Listener
	srv *grpc.Server
}

func (s *Server) Run() error {
	log.Println("start server")
	return s.srv.Serve(s.lis)
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}

func NewServer(
	lis net.Listener,
	repo model.Repository,
) *Server {
	srv := grpc.NewServer()
	app := application.NewSample(repo)

	sample.RegisterSampleServer(srv, app)
	reflection.Register(srv)

	return &Server{
		lis: lis,
		srv: srv,
	}
}
