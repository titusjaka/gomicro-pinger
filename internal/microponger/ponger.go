package microponger

import (
	"context"
	"log/slog"

	pb "github.com/titusjaka/gomicro-pinger/proto"
)

type Ponger struct{}

func (p *Ponger) Ping(_ context.Context, _ *pb.PingRequest, response *pb.PingResponse) error {
	slog.Info("Ponger.Pong")

	response.Message = "pong"
	return nil
}
