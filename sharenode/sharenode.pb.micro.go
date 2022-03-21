// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sharenode.proto

package sharenode

import (
	_ "./common"
	_ "./rescode"
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

// Api Endpoints for SharenodeService service

func NewSharenodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SharenodeService service

type SharenodeService interface {
	// 添加共享记录
	AddShare(ctx context.Context, in *AddShareReq, opts ...client.CallOption) (*AddShareRes, error)
	// 更新共享记录
	UpdateShare(ctx context.Context, in *UpdateShareReq, opts ...client.CallOption) (*UpdateShareRes, error)
	// 获取共享记录
	QueryShare(ctx context.Context, in *QueryShareReq, opts ...client.CallOption) (*QueryShareRes, error)
	// 更新模板
	UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, opts ...client.CallOption) (*UpdateTemplateRes, error)
	// 索引模板
	QueryTemplate(ctx context.Context, in *QueryTemplateReq, opts ...client.CallOption) (*QueryTemplateRes, error)
}

type sharenodeService struct {
	c    client.Client
	name string
}

func NewSharenodeService(name string, c client.Client) SharenodeService {
	return &sharenodeService{
		c:    c,
		name: name,
	}
}

func (c *sharenodeService) AddShare(ctx context.Context, in *AddShareReq, opts ...client.CallOption) (*AddShareRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.AddShare", in)
	out := new(AddShareRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharenodeService) UpdateShare(ctx context.Context, in *UpdateShareReq, opts ...client.CallOption) (*UpdateShareRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.UpdateShare", in)
	out := new(UpdateShareRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharenodeService) QueryShare(ctx context.Context, in *QueryShareReq, opts ...client.CallOption) (*QueryShareRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.QueryShare", in)
	out := new(QueryShareRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharenodeService) UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, opts ...client.CallOption) (*UpdateTemplateRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.UpdateTemplate", in)
	out := new(UpdateTemplateRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharenodeService) QueryTemplate(ctx context.Context, in *QueryTemplateReq, opts ...client.CallOption) (*QueryTemplateRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.QueryTemplate", in)
	out := new(QueryTemplateRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SharenodeService service

type SharenodeServiceHandler interface {
	// 添加共享记录
	AddShare(context.Context, *AddShareReq, *AddShareRes) error
	// 更新共享记录
	UpdateShare(context.Context, *UpdateShareReq, *UpdateShareRes) error
	// 获取共享记录
	QueryShare(context.Context, *QueryShareReq, *QueryShareRes) error
	// 更新模板
	UpdateTemplate(context.Context, *UpdateTemplateReq, *UpdateTemplateRes) error
	// 索引模板
	QueryTemplate(context.Context, *QueryTemplateReq, *QueryTemplateRes) error
}

func RegisterSharenodeServiceHandler(s server.Server, hdlr SharenodeServiceHandler, opts ...server.HandlerOption) error {
	type sharenodeService interface {
		AddShare(ctx context.Context, in *AddShareReq, out *AddShareRes) error
		UpdateShare(ctx context.Context, in *UpdateShareReq, out *UpdateShareRes) error
		QueryShare(ctx context.Context, in *QueryShareReq, out *QueryShareRes) error
		UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, out *UpdateTemplateRes) error
		QueryTemplate(ctx context.Context, in *QueryTemplateReq, out *QueryTemplateRes) error
	}
	type SharenodeService struct {
		sharenodeService
	}
	h := &sharenodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SharenodeService{h}, opts...))
}

type sharenodeServiceHandler struct {
	SharenodeServiceHandler
}

func (h *sharenodeServiceHandler) AddShare(ctx context.Context, in *AddShareReq, out *AddShareRes) error {
	return h.SharenodeServiceHandler.AddShare(ctx, in, out)
}

func (h *sharenodeServiceHandler) UpdateShare(ctx context.Context, in *UpdateShareReq, out *UpdateShareRes) error {
	return h.SharenodeServiceHandler.UpdateShare(ctx, in, out)
}

func (h *sharenodeServiceHandler) QueryShare(ctx context.Context, in *QueryShareReq, out *QueryShareRes) error {
	return h.SharenodeServiceHandler.QueryShare(ctx, in, out)
}

func (h *sharenodeServiceHandler) UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, out *UpdateTemplateRes) error {
	return h.SharenodeServiceHandler.UpdateTemplate(ctx, in, out)
}

func (h *sharenodeServiceHandler) QueryTemplate(ctx context.Context, in *QueryTemplateReq, out *QueryTemplateRes) error {
	return h.SharenodeServiceHandler.QueryTemplate(ctx, in, out)
}
