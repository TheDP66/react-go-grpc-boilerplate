package gapi

import (
	"github.com/TheDP66/react-go-grpc-server/pb"
	"github.com/TheDP66/react-go-grpc-server/util"
)

type Server struct {
	pb.UnimplementedReactGoServer
	config util.Config
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}
