package grpcponger

import (
	"context"
	"log/slog"

	pb "github.com/titusjaka/gomicro-pinger/proto"
)

type Ponger struct{}

func (p *Ponger) Ping(_ context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {
	slog.Info("Ponger.Pong")

	return &pb.PingResponse{
		Message: "pong",
	}, nil
}
