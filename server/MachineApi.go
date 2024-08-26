package server

import (
	"context"

	"github.com/becker63/tricky/pb"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedTrickyServer
}

func (*server) Start(
	ctx context.Context,
	in *pb.StartStopReq,
) (*pb.StartStopReq, error) {

	return nil, nil
}
