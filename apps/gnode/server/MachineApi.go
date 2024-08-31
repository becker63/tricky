package server

import (
	"context"

	"github.com/becker63/tricky/machines/Demo"
	"github.com/becker63/tricky/pb"
	"github.com/becker63/tricky/types"
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

	Demo.ConstructDemo(types.CommonBox{}).Start_i()

	return nil, nil
}
