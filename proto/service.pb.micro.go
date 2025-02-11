// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/service.proto

package gomicroPinger

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Pinger service

func NewPingerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Pinger service

type PingerService interface {
	Ping(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PingResponse, error)
}

type pingerService struct {
	c    client.Client
	name string
}

func NewPingerService(name string, c client.Client) PingerService {
	return &pingerService{
		c:    c,
		name: name,
	}
}

func (c *pingerService) Ping(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PingResponse, error) {
	req := c.c.NewRequest(c.name, "Pinger.Ping", in)
	out := new(PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pinger service

type PingerHandler interface {
	Ping(context.Context, *PingRequest, *PingResponse) error
}

func RegisterPingerHandler(s server.Server, hdlr PingerHandler, opts ...server.HandlerOption) error {
	type pinger interface {
		Ping(ctx context.Context, in *PingRequest, out *PingResponse) error
	}
	type Pinger struct {
		pinger
	}
	h := &pingerHandler{hdlr}
	return s.Handle(s.NewHandler(&Pinger{h}, opts...))
}

type pingerHandler struct {
	PingerHandler
}

func (h *pingerHandler) Ping(ctx context.Context, in *PingRequest, out *PingResponse) error {
	return h.PingerHandler.Ping(ctx, in, out)
}
