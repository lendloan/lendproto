// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cloudnode.proto

package cloudnode

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

// Api Endpoints for CloudnodeService service

func NewCloudnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CloudnodeService service

type CloudnodeService interface {
	// 获取目录
	Dir(ctx context.Context, in *DirReq, opts ...client.CallOption) (*DirRes, error)
	// 添加目录
	AddDir(ctx context.Context, in *AddDirReq, opts ...client.CallOption) (*AddDirRes, error)
	// 添加文件
	AddFile(ctx context.Context, in *AddFileReq, opts ...client.CallOption) (*AddFileRes, error)
	// 云存储属性
	Attr(ctx context.Context, in *AttrReq, opts ...client.CallOption) (*AttrRes, error)
	// 获取缩略图
	Thumbnail(ctx context.Context, in *ThumbnailReq, opts ...client.CallOption) (*ThumbnailRes, error)
}

type cloudnodeService struct {
	c    client.Client
	name string
}

func NewCloudnodeService(name string, c client.Client) CloudnodeService {
	return &cloudnodeService{
		c:    c,
		name: name,
	}
}

func (c *cloudnodeService) Dir(ctx context.Context, in *DirReq, opts ...client.CallOption) (*DirRes, error) {
	req := c.c.NewRequest(c.name, "CloudnodeService.Dir", in)
	out := new(DirRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudnodeService) AddDir(ctx context.Context, in *AddDirReq, opts ...client.CallOption) (*AddDirRes, error) {
	req := c.c.NewRequest(c.name, "CloudnodeService.AddDir", in)
	out := new(AddDirRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudnodeService) AddFile(ctx context.Context, in *AddFileReq, opts ...client.CallOption) (*AddFileRes, error) {
	req := c.c.NewRequest(c.name, "CloudnodeService.AddFile", in)
	out := new(AddFileRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudnodeService) Attr(ctx context.Context, in *AttrReq, opts ...client.CallOption) (*AttrRes, error) {
	req := c.c.NewRequest(c.name, "CloudnodeService.Attr", in)
	out := new(AttrRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudnodeService) Thumbnail(ctx context.Context, in *ThumbnailReq, opts ...client.CallOption) (*ThumbnailRes, error) {
	req := c.c.NewRequest(c.name, "CloudnodeService.Thumbnail", in)
	out := new(ThumbnailRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudnodeService service

type CloudnodeServiceHandler interface {
	// 获取目录
	Dir(context.Context, *DirReq, *DirRes) error
	// 添加目录
	AddDir(context.Context, *AddDirReq, *AddDirRes) error
	// 添加文件
	AddFile(context.Context, *AddFileReq, *AddFileRes) error
	// 云存储属性
	Attr(context.Context, *AttrReq, *AttrRes) error
	// 获取缩略图
	Thumbnail(context.Context, *ThumbnailReq, *ThumbnailRes) error
}

func RegisterCloudnodeServiceHandler(s server.Server, hdlr CloudnodeServiceHandler, opts ...server.HandlerOption) error {
	type cloudnodeService interface {
		Dir(ctx context.Context, in *DirReq, out *DirRes) error
		AddDir(ctx context.Context, in *AddDirReq, out *AddDirRes) error
		AddFile(ctx context.Context, in *AddFileReq, out *AddFileRes) error
		Attr(ctx context.Context, in *AttrReq, out *AttrRes) error
		Thumbnail(ctx context.Context, in *ThumbnailReq, out *ThumbnailRes) error
	}
	type CloudnodeService struct {
		cloudnodeService
	}
	h := &cloudnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CloudnodeService{h}, opts...))
}

type cloudnodeServiceHandler struct {
	CloudnodeServiceHandler
}

func (h *cloudnodeServiceHandler) Dir(ctx context.Context, in *DirReq, out *DirRes) error {
	return h.CloudnodeServiceHandler.Dir(ctx, in, out)
}

func (h *cloudnodeServiceHandler) AddDir(ctx context.Context, in *AddDirReq, out *AddDirRes) error {
	return h.CloudnodeServiceHandler.AddDir(ctx, in, out)
}

func (h *cloudnodeServiceHandler) AddFile(ctx context.Context, in *AddFileReq, out *AddFileRes) error {
	return h.CloudnodeServiceHandler.AddFile(ctx, in, out)
}

func (h *cloudnodeServiceHandler) Attr(ctx context.Context, in *AttrReq, out *AttrRes) error {
	return h.CloudnodeServiceHandler.Attr(ctx, in, out)
}

func (h *cloudnodeServiceHandler) Thumbnail(ctx context.Context, in *ThumbnailReq, out *ThumbnailRes) error {
	return h.CloudnodeServiceHandler.Thumbnail(ctx, in, out)
}
