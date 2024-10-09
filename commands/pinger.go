package commands

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/titusjaka/gomicro-pinger/grpcpinger"
	"github.com/titusjaka/gomicro-pinger/micropinger"
	pb "github.com/titusjaka/gomicro-pinger/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PingerCmd struct {
	Micro PingerMicroCmd `kong:"cmd,name='micro'"`
	GRPC  PingerGRPCCmd  `kong:"cmd,name='grpc'"`

	MicroGRPC PingerMicroGRPCCmd `kong:"cmd,name='micro-grpc'"`
}

type PingerMicroCmd struct {
	Flags Flags `kong:"embed"`
}

func (c PingerMicroCmd) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	mService := NewMicroService(ctx, c.Flags)

	gClient := pb.NewPingerService(
		microServiceName,
		mService.Client(),
	)

	mPinger := micropinger.NewPinger(gClient)

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := mPinger.Ping(ctx); err != nil {
				slog.With("err", err).Error("mPinger.Ping")
			}
		}
	}
}

type PingerGRPCCmd struct {
	Flags Flags `kong:"embed"`
}

func (c PingerGRPCCmd) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conn, err := grpc.NewClient(c.Flags.Listen(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}
	defer func(conn *grpc.ClientConn) {
		closeErr := conn.Close()
		if closeErr != nil {
			slog.With("err", closeErr).Error("conn.Close")
		}
	}(conn)

	client := pb.NewPingerClient(conn)
	pinger := grpcpinger.NewPinger(client)

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err = pinger.Ping(ctx); err != nil {
				slog.With("err", err).Error("mPinger.Ping")
			}
		}
	}
}

type PingerMicroGRPCCmd struct {
	Flags Flags `kong:"embed"`
}

func (c PingerMicroGRPCCmd) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	mService := NewMicroGRPCService(ctx, c.Flags)

	gClient := pb.NewPingerService(
		microServiceName,
		mService.Client(),
	)

	mPinger := micropinger.NewPinger(gClient)

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := mPinger.Ping(ctx); err != nil {
				slog.With("err", err).Error("mPinger.Ping")
			}
		}
	}
}
