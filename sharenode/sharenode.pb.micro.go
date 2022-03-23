// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sharenode.proto

package sharenode

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

// Api Endpoints for SharenodeService service

func NewSharenodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SharenodeService service

type SharenodeService interface {
	// 添加共享记录
	AddShare(ctx context.Context, in *AddShareReq, opts ...client.CallOption) (*AddShareRes, error)
	// 获取共享记录
	QueryShare(ctx context.Context, in *QueryShareReq, opts ...client.CallOption) (*QueryShareRes, error)
	// 更新模板
	UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, opts ...client.CallOption) (*UpdateTemplateRes, error)
	// 索引模板
	QueryTemplate(ctx context.Context, in *QueryTemplateReq, opts ...client.CallOption) (*QueryTemplateRes, error)
	// 获取模板数量
	TemplateCount(ctx context.Context, in *TemplateCountReq, opts ...client.CallOption) (*TemplateCountRes, error)
	// 更新share
	ShareTitle(ctx context.Context, in *ShareTitleReq, opts ...client.CallOption) (*ShareTitleRes, error)
	ShareMedia(ctx context.Context, in *ShareMediaReq, opts ...client.CallOption) (*ShareMediaRes, error)
	ShareUids(ctx context.Context, in *ShareUidsReq, opts ...client.CallOption) (*ShareUidsRes, error)
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

func (c *sharenodeService) TemplateCount(ctx context.Context, in *TemplateCountReq, opts ...client.CallOption) (*TemplateCountRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.TemplateCount", in)
	out := new(TemplateCountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharenodeService) ShareTitle(ctx context.Context, in *ShareTitleReq, opts ...client.CallOption) (*ShareTitleRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.ShareTitle", in)
	out := new(ShareTitleRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharenodeService) ShareMedia(ctx context.Context, in *ShareMediaReq, opts ...client.CallOption) (*ShareMediaRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.ShareMedia", in)
	out := new(ShareMediaRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharenodeService) ShareUids(ctx context.Context, in *ShareUidsReq, opts ...client.CallOption) (*ShareUidsRes, error) {
	req := c.c.NewRequest(c.name, "SharenodeService.ShareUids", in)
	out := new(ShareUidsRes)
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
	// 获取共享记录
	QueryShare(context.Context, *QueryShareReq, *QueryShareRes) error
	// 更新模板
	UpdateTemplate(context.Context, *UpdateTemplateReq, *UpdateTemplateRes) error
	// 索引模板
	QueryTemplate(context.Context, *QueryTemplateReq, *QueryTemplateRes) error
	// 获取模板数量
	TemplateCount(context.Context, *TemplateCountReq, *TemplateCountRes) error
	// 更新share
	ShareTitle(context.Context, *ShareTitleReq, *ShareTitleRes) error
	ShareMedia(context.Context, *ShareMediaReq, *ShareMediaRes) error
	ShareUids(context.Context, *ShareUidsReq, *ShareUidsRes) error
}

func RegisterSharenodeServiceHandler(s server.Server, hdlr SharenodeServiceHandler, opts ...server.HandlerOption) error {
	type sharenodeService interface {
		AddShare(ctx context.Context, in *AddShareReq, out *AddShareRes) error
		QueryShare(ctx context.Context, in *QueryShareReq, out *QueryShareRes) error
		UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, out *UpdateTemplateRes) error
		QueryTemplate(ctx context.Context, in *QueryTemplateReq, out *QueryTemplateRes) error
		TemplateCount(ctx context.Context, in *TemplateCountReq, out *TemplateCountRes) error
		ShareTitle(ctx context.Context, in *ShareTitleReq, out *ShareTitleRes) error
		ShareMedia(ctx context.Context, in *ShareMediaReq, out *ShareMediaRes) error
		ShareUids(ctx context.Context, in *ShareUidsReq, out *ShareUidsRes) error
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

func (h *sharenodeServiceHandler) QueryShare(ctx context.Context, in *QueryShareReq, out *QueryShareRes) error {
	return h.SharenodeServiceHandler.QueryShare(ctx, in, out)
}

func (h *sharenodeServiceHandler) UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, out *UpdateTemplateRes) error {
	return h.SharenodeServiceHandler.UpdateTemplate(ctx, in, out)
}

func (h *sharenodeServiceHandler) QueryTemplate(ctx context.Context, in *QueryTemplateReq, out *QueryTemplateRes) error {
	return h.SharenodeServiceHandler.QueryTemplate(ctx, in, out)
}

func (h *sharenodeServiceHandler) TemplateCount(ctx context.Context, in *TemplateCountReq, out *TemplateCountRes) error {
	return h.SharenodeServiceHandler.TemplateCount(ctx, in, out)
}

func (h *sharenodeServiceHandler) ShareTitle(ctx context.Context, in *ShareTitleReq, out *ShareTitleRes) error {
	return h.SharenodeServiceHandler.ShareTitle(ctx, in, out)
}

func (h *sharenodeServiceHandler) ShareMedia(ctx context.Context, in *ShareMediaReq, out *ShareMediaRes) error {
	return h.SharenodeServiceHandler.ShareMedia(ctx, in, out)
}

func (h *sharenodeServiceHandler) ShareUids(ctx context.Context, in *ShareUidsReq, out *ShareUidsRes) error {
	return h.SharenodeServiceHandler.ShareUids(ctx, in, out)
}
