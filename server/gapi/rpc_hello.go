package gapi

import (
	"context"

	"github.com/TheDP66/react-go-grpc-server/pb"
)

func (server *Server) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	rsp := &pb.HelloResponse{
		Text: req.GetText(),
	}
	return rsp, nil
}
