package commands

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/titusjaka/gomicro-pinger/internal/grpcponger"
	"github.com/titusjaka/gomicro-pinger/internal/microponger"
	pb "github.com/titusjaka/gomicro-pinger/proto"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type PongerCmd struct {
	Micro PongerMicroCmd `kong:"cmd,name='micro'"`
	GRPC  PongerGRPCCmd  `kong:"cmd,name='grpc'"`

	MicroGRPC PongerMicroGRPCCmd `kong:"cmd,name='micro-grpc'"`
}

type PongerMicroCmd struct {
	Flags Flags `kong:"embed"`
}

func (c PongerMicroCmd) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	mService := NewMicroService(ctx, c.Flags)

	if err := pb.RegisterPingerHandler(mService.Server(), &microponger.Ponger{}); err != nil {
		return fmt.Errorf("register handler: %w", err)
	}

	slog.Info("starting micro server")

	return mService.Run()
}

type PongerGRPCCmd struct {
	Flags Flags `kong:"embed"`
}

func (c PongerGRPCCmd) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	gr, ctx := errgroup.WithContext(ctx)

	lister, err := net.Listen("tcp", c.Flags.Listen())
	if err != nil {
		return fmt.Errorf("create listener: %w", err)
	}

	pongerServer := grpcponger.Ponger{}

	grpcServer := grpc.NewServer()
	pb.RegisterPingerServer(grpcServer, &pongerServer)

	gr.Go(func() error {
		<-ctx.Done()
		grpcServer.GracefulStop()
		return nil
	})

	gr.Go(func() error {
		slog.Info("starting grpc server")
		return grpcServer.Serve(lister)
	})

	return gr.Wait()
}

type PongerMicroGRPCCmd struct {
	Flags Flags `kong:"embed"`
}

func (c PongerMicroGRPCCmd) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	mService := NewMicroGRPCService(ctx, c.Flags)

	if err := pb.RegisterPingerHandler(mService.Server(), &microponger.Ponger{}); err != nil {
		return fmt.Errorf("register handler: %w", err)
	}

	slog.Info("starting micro-grpc server")

	return mService.Run()
}
