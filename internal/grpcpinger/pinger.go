package grpcpinger

import (
	"context"
	"log/slog"

	pb "github.com/titusjaka/gomicro-pinger/proto"
)

type Pinger struct {
	client pb.PingerClient
}

func NewPinger(client pb.PingerClient) *Pinger {
	return &Pinger{client: client}
}

func (p *Pinger) Ping(ctx context.Context) error {
	slog.Info("Pinger.Ping")

	_, err := p.client.Ping(ctx, &pb.PingRequest{
		Message: "ping",
	})

	return err
}
