// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: goodsnode.proto

package goodsnode

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

// Api Endpoints for GoodsnodeService service

func NewGoodsnodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GoodsnodeService service

type GoodsnodeService interface {
	// 我的物品
	AddGoodsCategory(ctx context.Context, in *AddGoodsCategoryReq, opts ...client.CallOption) (*AddGoodsCategoryRes, error)
	GoodsCategory(ctx context.Context, in *GoodsCategoryReq, opts ...client.CallOption) (*GoodsCategoryRes, error)
	AddGoodsTotal(ctx context.Context, in *AddGoodsTotalReq, opts ...client.CallOption) (*AddGoodsTotalRes, error)
	UpdateGoodsCate(ctx context.Context, in *UpdateGoodsCateReq, opts ...client.CallOption) (*UpdateGoodsCateRes, error)
	UpdateGoodsName(ctx context.Context, in *UpdateGoodsNameReq, opts ...client.CallOption) (*UpdateGoodsNameRes, error)
	UpdateGoodsIntro(ctx context.Context, in *UpdateGoodsIntroReq, opts ...client.CallOption) (*UpdateGoodsIntroRes, error)
	UpdateGoodsPrice(ctx context.Context, in *UpdateGoodsPriceReq, opts ...client.CallOption) (*UpdateGoodsPriceRes, error)
	Goods(ctx context.Context, in *GoodsReq, opts ...client.CallOption) (*GoodsRes, error)
}

type goodsnodeService struct {
	c    client.Client
	name string
}

func NewGoodsnodeService(name string, c client.Client) GoodsnodeService {
	return &goodsnodeService{
		c:    c,
		name: name,
	}
}

func (c *goodsnodeService) AddGoodsCategory(ctx context.Context, in *AddGoodsCategoryReq, opts ...client.CallOption) (*AddGoodsCategoryRes, error) {
	req := c.c.NewRequest(c.name, "GoodsnodeService.AddGoodsCategory", in)
	out := new(AddGoodsCategoryRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsnodeService) GoodsCategory(ctx context.Context, in *GoodsCategoryReq, opts ...client.CallOption) (*GoodsCategoryRes, error) {
	req := c.c.NewRequest(c.name, "GoodsnodeService.GoodsCategory", in)
	out := new(GoodsCategoryRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsnodeService) AddGoodsTotal(ctx context.Context, in *AddGoodsTotalReq, opts ...client.CallOption) (*AddGoodsTotalRes, error) {
	req := c.c.NewRequest(c.name, "GoodsnodeService.AddGoodsTotal", in)
	out := new(AddGoodsTotalRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsnodeService) UpdateGoodsCate(ctx context.Context, in *UpdateGoodsCateReq, opts ...client.CallOption) (*UpdateGoodsCateRes, error) {
	req := c.c.NewRequest(c.name, "GoodsnodeService.UpdateGoodsCate", in)
	out := new(UpdateGoodsCateRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsnodeService) UpdateGoodsName(ctx context.Context, in *UpdateGoodsNameReq, opts ...client.CallOption) (*UpdateGoodsNameRes, error) {
	req := c.c.NewRequest(c.name, "GoodsnodeService.UpdateGoodsName", in)
	out := new(UpdateGoodsNameRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsnodeService) UpdateGoodsIntro(ctx context.Context, in *UpdateGoodsIntroReq, opts ...client.CallOption) (*UpdateGoodsIntroRes, error) {
	req := c.c.NewRequest(c.name, "GoodsnodeService.UpdateGoodsIntro", in)
	out := new(UpdateGoodsIntroRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsnodeService) UpdateGoodsPrice(ctx context.Context, in *UpdateGoodsPriceReq, opts ...client.CallOption) (*UpdateGoodsPriceRes, error) {
	req := c.c.NewRequest(c.name, "GoodsnodeService.UpdateGoodsPrice", in)
	out := new(UpdateGoodsPriceRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsnodeService) Goods(ctx context.Context, in *GoodsReq, opts ...client.CallOption) (*GoodsRes, error) {
	req := c.c.NewRequest(c.name, "GoodsnodeService.Goods", in)
	out := new(GoodsRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoodsnodeService service

type GoodsnodeServiceHandler interface {
	// 我的物品
	AddGoodsCategory(context.Context, *AddGoodsCategoryReq, *AddGoodsCategoryRes) error
	GoodsCategory(context.Context, *GoodsCategoryReq, *GoodsCategoryRes) error
	AddGoodsTotal(context.Context, *AddGoodsTotalReq, *AddGoodsTotalRes) error
	UpdateGoodsCate(context.Context, *UpdateGoodsCateReq, *UpdateGoodsCateRes) error
	UpdateGoodsName(context.Context, *UpdateGoodsNameReq, *UpdateGoodsNameRes) error
	UpdateGoodsIntro(context.Context, *UpdateGoodsIntroReq, *UpdateGoodsIntroRes) error
	UpdateGoodsPrice(context.Context, *UpdateGoodsPriceReq, *UpdateGoodsPriceRes) error
	Goods(context.Context, *GoodsReq, *GoodsRes) error
}

func RegisterGoodsnodeServiceHandler(s server.Server, hdlr GoodsnodeServiceHandler, opts ...server.HandlerOption) error {
	type goodsnodeService interface {
		AddGoodsCategory(ctx context.Context, in *AddGoodsCategoryReq, out *AddGoodsCategoryRes) error
		GoodsCategory(ctx context.Context, in *GoodsCategoryReq, out *GoodsCategoryRes) error
		AddGoodsTotal(ctx context.Context, in *AddGoodsTotalReq, out *AddGoodsTotalRes) error
		UpdateGoodsCate(ctx context.Context, in *UpdateGoodsCateReq, out *UpdateGoodsCateRes) error
		UpdateGoodsName(ctx context.Context, in *UpdateGoodsNameReq, out *UpdateGoodsNameRes) error
		UpdateGoodsIntro(ctx context.Context, in *UpdateGoodsIntroReq, out *UpdateGoodsIntroRes) error
		UpdateGoodsPrice(ctx context.Context, in *UpdateGoodsPriceReq, out *UpdateGoodsPriceRes) error
		Goods(ctx context.Context, in *GoodsReq, out *GoodsRes) error
	}
	type GoodsnodeService struct {
		goodsnodeService
	}
	h := &goodsnodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GoodsnodeService{h}, opts...))
}

type goodsnodeServiceHandler struct {
	GoodsnodeServiceHandler
}

func (h *goodsnodeServiceHandler) AddGoodsCategory(ctx context.Context, in *AddGoodsCategoryReq, out *AddGoodsCategoryRes) error {
	return h.GoodsnodeServiceHandler.AddGoodsCategory(ctx, in, out)
}

func (h *goodsnodeServiceHandler) GoodsCategory(ctx context.Context, in *GoodsCategoryReq, out *GoodsCategoryRes) error {
	return h.GoodsnodeServiceHandler.GoodsCategory(ctx, in, out)
}

func (h *goodsnodeServiceHandler) AddGoodsTotal(ctx context.Context, in *AddGoodsTotalReq, out *AddGoodsTotalRes) error {
	return h.GoodsnodeServiceHandler.AddGoodsTotal(ctx, in, out)
}

func (h *goodsnodeServiceHandler) UpdateGoodsCate(ctx context.Context, in *UpdateGoodsCateReq, out *UpdateGoodsCateRes) error {
	return h.GoodsnodeServiceHandler.UpdateGoodsCate(ctx, in, out)
}

func (h *goodsnodeServiceHandler) UpdateGoodsName(ctx context.Context, in *UpdateGoodsNameReq, out *UpdateGoodsNameRes) error {
	return h.GoodsnodeServiceHandler.UpdateGoodsName(ctx, in, out)
}

func (h *goodsnodeServiceHandler) UpdateGoodsIntro(ctx context.Context, in *UpdateGoodsIntroReq, out *UpdateGoodsIntroRes) error {
	return h.GoodsnodeServiceHandler.UpdateGoodsIntro(ctx, in, out)
}

func (h *goodsnodeServiceHandler) UpdateGoodsPrice(ctx context.Context, in *UpdateGoodsPriceReq, out *UpdateGoodsPriceRes) error {
	return h.GoodsnodeServiceHandler.UpdateGoodsPrice(ctx, in, out)
}

func (h *goodsnodeServiceHandler) Goods(ctx context.Context, in *GoodsReq, out *GoodsRes) error {
	return h.GoodsnodeServiceHandler.Goods(ctx, in, out)
}
