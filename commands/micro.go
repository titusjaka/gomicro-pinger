package commands

import (
	"context"
	"log/slog"
	"os"

	mcg "github.com/go-micro/plugins/v4/client/grpc"
	goMicroK8s "github.com/go-micro/plugins/v4/registry/kubernetes"
	msg "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

const (
	microServiceName = "gomicroPinger"
)

func NewMicroService(ctx context.Context, flags Flags) micro.Service {
	mService := micro.NewService(
		micro.Server(NewMicroServer(ctx)),
		micro.Logger(logger.NewLogger()),
		micro.Name(microServiceName),
		micro.Address(flags.Listen()),
		micro.Registry(NewMicroRegistry()),
	)

	return mService
}

func NewMicroGRPCService(ctx context.Context, flags Flags) micro.Service {
	mService := micro.NewService(
		micro.Server(NewMicroGRPCServer(ctx)),
		micro.Client(NewMicroGRPCClient()),
		micro.Logger(logger.NewLogger()),
		micro.Name(microServiceName),
		micro.Address(flags.Listen()),
		micro.Registry(NewMicroRegistry()),
	)

	return mService
}

func NewMicroServer(ctx context.Context) server.Server {
	return server.NewServer(
		server.Context(ctx),
		server.WithLogger(logger.NewLogger()),
	)
}

func NewMicroGRPCServer(ctx context.Context) server.Server {
	return msg.NewServer(
		server.Context(ctx),
		server.WithLogger(logger.NewLogger()),
	)
}

func NewMicroGRPCClient() client.Client {
	return mcg.NewClient()
}

func NewMicroRegistry() registry.Registry {
	instance := registry.DefaultRegistry
	// checking if it's kubernetes pod
	if len(os.Getenv("KUBERNETES_SERVICE_HOST")) > 0 {
		slog.Info("Using kubernetes registry")
		instance = goMicroK8s.NewRegistry()
	} else {
		slog.Info("Using default registry")
	}
	return instance
}
