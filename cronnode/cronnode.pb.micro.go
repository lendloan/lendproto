// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cronnode.proto

package cronnode

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

// Api Endpoints for CronnodeService service

func NewCronnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CronnodeService service

type CronnodeService interface {
	// 提交实名
	CronCall(ctx context.Context, in *CronCallReq, opts ...client.CallOption) (*CronCallRes, error)
	// 添加任务
	AddTask(ctx context.Context, in *AddTaskReq, opts ...client.CallOption) (*AddTaskRes, error)
	// 获取任务
	GetTask(ctx context.Context, in *GetTaskReq, opts ...client.CallOption) (*GetTaskRes, error)
	// 获取任务数量
	GetTaskCount(ctx context.Context, in *GetTaskCountReq, opts ...client.CallOption) (*GetTaskCountRes, error)
}

type cronnodeService struct {
	c    client.Client
	name string
}

func NewCronnodeService(name string, c client.Client) CronnodeService {
	return &cronnodeService{
		c:    c,
		name: name,
	}
}

func (c *cronnodeService) CronCall(ctx context.Context, in *CronCallReq, opts ...client.CallOption) (*CronCallRes, error) {
	req := c.c.NewRequest(c.name, "CronnodeService.CronCall", in)
	out := new(CronCallRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronnodeService) AddTask(ctx context.Context, in *AddTaskReq, opts ...client.CallOption) (*AddTaskRes, error) {
	req := c.c.NewRequest(c.name, "CronnodeService.AddTask", in)
	out := new(AddTaskRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronnodeService) GetTask(ctx context.Context, in *GetTaskReq, opts ...client.CallOption) (*GetTaskRes, error) {
	req := c.c.NewRequest(c.name, "CronnodeService.GetTask", in)
	out := new(GetTaskRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronnodeService) GetTaskCount(ctx context.Context, in *GetTaskCountReq, opts ...client.CallOption) (*GetTaskCountRes, error) {
	req := c.c.NewRequest(c.name, "CronnodeService.GetTaskCount", in)
	out := new(GetTaskCountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CronnodeService service

type CronnodeServiceHandler interface {
	// 提交实名
	CronCall(context.Context, *CronCallReq, *CronCallRes) error
	// 添加任务
	AddTask(context.Context, *AddTaskReq, *AddTaskRes) error
	// 获取任务
	GetTask(context.Context, *GetTaskReq, *GetTaskRes) error
	// 获取任务数量
	GetTaskCount(context.Context, *GetTaskCountReq, *GetTaskCountRes) error
}

func RegisterCronnodeServiceHandler(s server.Server, hdlr CronnodeServiceHandler, opts ...server.HandlerOption) error {
	type cronnodeService interface {
		CronCall(ctx context.Context, in *CronCallReq, out *CronCallRes) error
		AddTask(ctx context.Context, in *AddTaskReq, out *AddTaskRes) error
		GetTask(ctx context.Context, in *GetTaskReq, out *GetTaskRes) error
		GetTaskCount(ctx context.Context, in *GetTaskCountReq, out *GetTaskCountRes) error
	}
	type CronnodeService struct {
		cronnodeService
	}
	h := &cronnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CronnodeService{h}, opts...))
}

type cronnodeServiceHandler struct {
	CronnodeServiceHandler
}

func (h *cronnodeServiceHandler) CronCall(ctx context.Context, in *CronCallReq, out *CronCallRes) error {
	return h.CronnodeServiceHandler.CronCall(ctx, in, out)
}

func (h *cronnodeServiceHandler) AddTask(ctx context.Context, in *AddTaskReq, out *AddTaskRes) error {
	return h.CronnodeServiceHandler.AddTask(ctx, in, out)
}

func (h *cronnodeServiceHandler) GetTask(ctx context.Context, in *GetTaskReq, out *GetTaskRes) error {
	return h.CronnodeServiceHandler.GetTask(ctx, in, out)
}

func (h *cronnodeServiceHandler) GetTaskCount(ctx context.Context, in *GetTaskCountReq, out *GetTaskCountRes) error {
	return h.CronnodeServiceHandler.GetTaskCount(ctx, in, out)
}
