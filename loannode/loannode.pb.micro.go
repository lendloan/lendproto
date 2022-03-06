// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: loannode.proto

package loannode

import (
	_ "github.com/lendloan/lendproto/common"
	_ "github.com/lendloan/lendproto/rescode"
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

// Api Endpoints for LoannodeService service

func NewLoannodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LoannodeService service

type LoannodeService interface {
}

type loannodeService struct {
	c    client.Client
	name string
}

func NewLoannodeService(name string, c client.Client) LoannodeService {
	return &loannodeService{
		c:    c,
		name: name,
	}
}

// Server API for LoannodeService service

type LoannodeServiceHandler interface {
}

func RegisterLoannodeServiceHandler(s server.Server, hdlr LoannodeServiceHandler, opts ...server.HandlerOption) error {
	type loannodeService interface {
	}
	type LoannodeService struct {
		loannodeService
	}
	h := &loannodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LoannodeService{h}, opts...))
}

type loannodeServiceHandler struct {
	LoannodeServiceHandler
}
