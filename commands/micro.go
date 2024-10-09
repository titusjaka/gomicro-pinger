package commands

import (
	"context"

	mcg "github.com/go-micro/plugins/v4/client/grpc"
	msg "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
)

const (
	microServiceName = "gomicroPinger"
)

func NewMicroService(ctx context.Context, flags Flags) micro.Service {
	mService := micro.NewService(
		micro.Server(server.NewServer(
			server.Context(ctx),
			server.WithLogger(logger.NewLogger()),
		)),
		micro.Logger(logger.NewLogger()),
		micro.Name(microServiceName),
		micro.Address(flags.Listen()),
	)

	return mService
}

func NewMicroGRPCService(ctx context.Context, flags Flags) micro.Service {
	mService := micro.NewService(
		micro.Server(msg.NewServer(
			server.Context(ctx),
			server.WithLogger(logger.NewLogger()),
		)),
		micro.Client(mcg.NewClient()),
		micro.Logger(logger.NewLogger()),
		micro.Name(microServiceName),
		micro.Address(flags.Listen()),
	)

	return mService
}
