// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: datanode.proto

package datanode

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

// Api Endpoints for DatanodeService service

func NewDatanodeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DatanodeService service

type DatanodeService interface {
	// ---------- 用户接口 ------- //
	// 创建新用户
	CreateUser(ctx context.Context, in *NewUserReq, opts ...client.CallOption) (*UserRes, error)
	// 获取用户信息
	UserInfo(ctx context.Context, in *UserInfoReq, opts ...client.CallOption) (*UserRes, error)
	// 搜索用户信息
	SearchUser(ctx context.Context, in *SearchUserReq, opts ...client.CallOption) (*SearchUserRes, error)
	// 更新用户信息
	UpdateUserInfo(ctx context.Context, in *UpdateUserinfoReq, opts ...client.CallOption) (*UpdateUserinfoRes, error)
	// ---------------- 文件 ------------------//
	// 获取目录,目录之间的/替换成-
	FileDirOne(ctx context.Context, in *FileDirReq, opts ...client.CallOption) (*FileDirRes, error)
	// 添加目录
	FileDirAdd(ctx context.Context, in *AddDirReq, opts ...client.CallOption) (*AddDirRes, error)
	// 添加文件
	AddFile(ctx context.Context, in *AddFileReq, opts ...client.CallOption) (*AddFileRes, error)
	// 获取云盘属性，容量和文件数量
	YunSaveAttr(ctx context.Context, in *FileAttrReq, opts ...client.CallOption) (*FileAttrRes, error)
	// 获取图像缩略图
	Thumbnail(ctx context.Context, in *ThumbnailReq, opts ...client.CallOption) (*ThumbnailRes, error)
	// ------------- key-value ----------------//
	// 设置缓存
	SetKeyValue(ctx context.Context, in *SetKeyvalueReq, opts ...client.CallOption) (*SetKeyvalueRes, error)
	// 获取缓存
	GetKeyValue(ctx context.Context, in *GetKeyvalueReq, opts ...client.CallOption) (*GetKeyvalueRes, error)
	// 删除
	DelKeyValue(ctx context.Context, in *DelKeyvalueReq, opts ...client.CallOption) (*DelKeyvalueRes, error)
	//用户信息
	SetUserInfo(ctx context.Context, in *SetUserInfoReq, opts ...client.CallOption) (*SetUserInfoRes, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...client.CallOption) (*GetUserInfoRes, error)
	//微信用户信息操作
	SetWechatUinfo(ctx context.Context, in *SetWechatUinfoReq, opts ...client.CallOption) (*SetWechatUinfoRes, error)
	GetWechatUinfo(ctx context.Context, in *GetWechatUinfoReq, opts ...client.CallOption) (*GetWechatUinfoRes, error)
	// 支付宝用户信息操作
	SetAlipayUinfo(ctx context.Context, in *SetAlipayUinfoReq, opts ...client.CallOption) (*SetAlipayUinfoRes, error)
	GetAlipayUinfo(ctx context.Context, in *GetAlipayUinfoReq, opts ...client.CallOption) (*GetAlipayUinfoRes, error)
	// 请求百度实体信息
	SetBaiduEntity(ctx context.Context, in *SetBaiduEntityReq, opts ...client.CallOption) (*SetBaiduEntityRes, error)
	BaiduEntity(ctx context.Context, in *BaiduEntityReq, opts ...client.CallOption) (*BaiduEntityRes, error)
	// 提交实名
	SubmitCert(ctx context.Context, in *SubmitCertReq, opts ...client.CallOption) (*SubmitCertRes, error)
	// 更新实名状态
	CertStatus(ctx context.Context, in *CertStatusReq, opts ...client.CallOption) (*CertStatusRes, error)
	// 获取实名信息
	CertInfo(ctx context.Context, in *CertInfoReq, opts ...client.CallOption) (*CertInfoRes, error)
	// 实名日志记录
	CertFlow(ctx context.Context, in *CertFlowReq, opts ...client.CallOption) (*CertFlowRes, error)
	// 添加日志
	AddUserLog(ctx context.Context, in *AddUserLogReq, opts ...client.CallOption) (*AddUserLogRes, error)
	// 获取日志
	QueryUserLog(ctx context.Context, in *QueryUserLogReq, opts ...client.CallOption) (*QueryUserLogRes, error)
	// 日志数量
	UserLogCount(ctx context.Context, in *UserLogCountReq, opts ...client.CallOption) (*UserLogCountRes, error)
	// 添加任务
	AddTask(ctx context.Context, in *AddTaskReq, opts ...client.CallOption) (*AddTaskRes, error)
	// 获取任务
	GetTask(ctx context.Context, in *GetTaskReq, opts ...client.CallOption) (*GetTaskRes, error)
	// 获取任务数量
	GetTaskCount(ctx context.Context, in *GetTaskCountReq, opts ...client.CallOption) (*GetTaskCountRes, error)
	// 添加包信息
	AddPkg(ctx context.Context, in *AddPkgReq, opts ...client.CallOption) (*AddPkgRes, error)
	// 删除包
	DelPkg(ctx context.Context, in *DelPkgReq, opts ...client.CallOption) (*DelPkgRes, error)
	// 获取包列表
	PkgLists(ctx context.Context, in *PkgListReq, opts ...client.CallOption) (*PkgListRes, error)
	// 获取包数量
	PkgCount(ctx context.Context, in *PkgCountReq, opts ...client.CallOption) (*PkgCountRes, error)
	// 获取vip对应的限制
	VipLimit(ctx context.Context, in *VipLimitReq, opts ...client.CallOption) (*VipLimitRes, error)
	// vip产品列表
	VipProduct(ctx context.Context, in *VipProductReq, opts ...client.CallOption) (*VipProductRes, error)
	// vip产品介绍
	VipDesc(ctx context.Context, in *VipDescReq, opts ...client.CallOption) (*VipDescRes, error)
	// vip充值
	VipPay(ctx context.Context, in *VipPayReq, opts ...client.CallOption) (*VipPayRes, error)
	// 更新vip充值状态
	VipOrderStatus(ctx context.Context, in *VipOrderStatusReq, opts ...client.CallOption) (*VipOrderStatusRes, error)
	// 获取vip充值列表
	VipOrderList(ctx context.Context, in *VipOrderListReq, opts ...client.CallOption) (*VipOrderListRes, error)
	// 用户积分操作
	UserScore(ctx context.Context, in *UserScoreReq, opts ...client.CallOption) (*UserScoreRes, error)
	// 用户vip操作
	UserVip(ctx context.Context, in *UserVipReq, opts ...client.CallOption) (*UserVipRes, error)
	RefreshUserCoin(ctx context.Context, in *RefreshUserCoinReq, opts ...client.CallOption) (*RefreshUserCoinRes, error)
	// 服务调用日志
	FootLog(ctx context.Context, in *FootLogReq, opts ...client.CallOption) (*FootLogRes, error)
}

type datanodeService struct {
	c    client.Client
	name string
}

func NewDatanodeService(name string, c client.Client) DatanodeService {
	return &datanodeService{
		c:    c,
		name: name,
	}
}

func (c *datanodeService) CreateUser(ctx context.Context, in *NewUserReq, opts ...client.CallOption) (*UserRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.CreateUser", in)
	out := new(UserRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) UserInfo(ctx context.Context, in *UserInfoReq, opts ...client.CallOption) (*UserRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.UserInfo", in)
	out := new(UserRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) SearchUser(ctx context.Context, in *SearchUserReq, opts ...client.CallOption) (*SearchUserRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.SearchUser", in)
	out := new(SearchUserRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) UpdateUserInfo(ctx context.Context, in *UpdateUserinfoReq, opts ...client.CallOption) (*UpdateUserinfoRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.UpdateUserInfo", in)
	out := new(UpdateUserinfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) FileDirOne(ctx context.Context, in *FileDirReq, opts ...client.CallOption) (*FileDirRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.FileDirOne", in)
	out := new(FileDirRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) FileDirAdd(ctx context.Context, in *AddDirReq, opts ...client.CallOption) (*AddDirRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.FileDirAdd", in)
	out := new(AddDirRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) AddFile(ctx context.Context, in *AddFileReq, opts ...client.CallOption) (*AddFileRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.AddFile", in)
	out := new(AddFileRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) YunSaveAttr(ctx context.Context, in *FileAttrReq, opts ...client.CallOption) (*FileAttrRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.YunSaveAttr", in)
	out := new(FileAttrRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) Thumbnail(ctx context.Context, in *ThumbnailReq, opts ...client.CallOption) (*ThumbnailRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.Thumbnail", in)
	out := new(ThumbnailRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) SetKeyValue(ctx context.Context, in *SetKeyvalueReq, opts ...client.CallOption) (*SetKeyvalueRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.SetKeyValue", in)
	out := new(SetKeyvalueRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) GetKeyValue(ctx context.Context, in *GetKeyvalueReq, opts ...client.CallOption) (*GetKeyvalueRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.GetKeyValue", in)
	out := new(GetKeyvalueRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) DelKeyValue(ctx context.Context, in *DelKeyvalueReq, opts ...client.CallOption) (*DelKeyvalueRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.DelKeyValue", in)
	out := new(DelKeyvalueRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) SetUserInfo(ctx context.Context, in *SetUserInfoReq, opts ...client.CallOption) (*SetUserInfoRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.SetUserInfo", in)
	out := new(SetUserInfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...client.CallOption) (*GetUserInfoRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.GetUserInfo", in)
	out := new(GetUserInfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) SetWechatUinfo(ctx context.Context, in *SetWechatUinfoReq, opts ...client.CallOption) (*SetWechatUinfoRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.SetWechatUinfo", in)
	out := new(SetWechatUinfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) GetWechatUinfo(ctx context.Context, in *GetWechatUinfoReq, opts ...client.CallOption) (*GetWechatUinfoRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.GetWechatUinfo", in)
	out := new(GetWechatUinfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) SetAlipayUinfo(ctx context.Context, in *SetAlipayUinfoReq, opts ...client.CallOption) (*SetAlipayUinfoRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.SetAlipayUinfo", in)
	out := new(SetAlipayUinfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) GetAlipayUinfo(ctx context.Context, in *GetAlipayUinfoReq, opts ...client.CallOption) (*GetAlipayUinfoRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.GetAlipayUinfo", in)
	out := new(GetAlipayUinfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) SetBaiduEntity(ctx context.Context, in *SetBaiduEntityReq, opts ...client.CallOption) (*SetBaiduEntityRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.SetBaiduEntity", in)
	out := new(SetBaiduEntityRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) BaiduEntity(ctx context.Context, in *BaiduEntityReq, opts ...client.CallOption) (*BaiduEntityRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.BaiduEntity", in)
	out := new(BaiduEntityRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) SubmitCert(ctx context.Context, in *SubmitCertReq, opts ...client.CallOption) (*SubmitCertRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.SubmitCert", in)
	out := new(SubmitCertRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) CertStatus(ctx context.Context, in *CertStatusReq, opts ...client.CallOption) (*CertStatusRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.CertStatus", in)
	out := new(CertStatusRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) CertInfo(ctx context.Context, in *CertInfoReq, opts ...client.CallOption) (*CertInfoRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.CertInfo", in)
	out := new(CertInfoRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) CertFlow(ctx context.Context, in *CertFlowReq, opts ...client.CallOption) (*CertFlowRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.CertFlow", in)
	out := new(CertFlowRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) AddUserLog(ctx context.Context, in *AddUserLogReq, opts ...client.CallOption) (*AddUserLogRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.AddUserLog", in)
	out := new(AddUserLogRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) QueryUserLog(ctx context.Context, in *QueryUserLogReq, opts ...client.CallOption) (*QueryUserLogRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.QueryUserLog", in)
	out := new(QueryUserLogRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) UserLogCount(ctx context.Context, in *UserLogCountReq, opts ...client.CallOption) (*UserLogCountRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.UserLogCount", in)
	out := new(UserLogCountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) AddTask(ctx context.Context, in *AddTaskReq, opts ...client.CallOption) (*AddTaskRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.AddTask", in)
	out := new(AddTaskRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) GetTask(ctx context.Context, in *GetTaskReq, opts ...client.CallOption) (*GetTaskRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.GetTask", in)
	out := new(GetTaskRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) GetTaskCount(ctx context.Context, in *GetTaskCountReq, opts ...client.CallOption) (*GetTaskCountRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.GetTaskCount", in)
	out := new(GetTaskCountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) AddPkg(ctx context.Context, in *AddPkgReq, opts ...client.CallOption) (*AddPkgRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.AddPkg", in)
	out := new(AddPkgRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) DelPkg(ctx context.Context, in *DelPkgReq, opts ...client.CallOption) (*DelPkgRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.DelPkg", in)
	out := new(DelPkgRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) PkgLists(ctx context.Context, in *PkgListReq, opts ...client.CallOption) (*PkgListRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.PkgLists", in)
	out := new(PkgListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) PkgCount(ctx context.Context, in *PkgCountReq, opts ...client.CallOption) (*PkgCountRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.PkgCount", in)
	out := new(PkgCountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) VipLimit(ctx context.Context, in *VipLimitReq, opts ...client.CallOption) (*VipLimitRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.VipLimit", in)
	out := new(VipLimitRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) VipProduct(ctx context.Context, in *VipProductReq, opts ...client.CallOption) (*VipProductRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.VipProduct", in)
	out := new(VipProductRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) VipDesc(ctx context.Context, in *VipDescReq, opts ...client.CallOption) (*VipDescRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.VipDesc", in)
	out := new(VipDescRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) VipPay(ctx context.Context, in *VipPayReq, opts ...client.CallOption) (*VipPayRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.VipPay", in)
	out := new(VipPayRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) VipOrderStatus(ctx context.Context, in *VipOrderStatusReq, opts ...client.CallOption) (*VipOrderStatusRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.VipOrderStatus", in)
	out := new(VipOrderStatusRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) VipOrderList(ctx context.Context, in *VipOrderListReq, opts ...client.CallOption) (*VipOrderListRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.VipOrderList", in)
	out := new(VipOrderListRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) UserScore(ctx context.Context, in *UserScoreReq, opts ...client.CallOption) (*UserScoreRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.UserScore", in)
	out := new(UserScoreRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) UserVip(ctx context.Context, in *UserVipReq, opts ...client.CallOption) (*UserVipRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.UserVip", in)
	out := new(UserVipRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) RefreshUserCoin(ctx context.Context, in *RefreshUserCoinReq, opts ...client.CallOption) (*RefreshUserCoinRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.RefreshUserCoin", in)
	out := new(RefreshUserCoinRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datanodeService) FootLog(ctx context.Context, in *FootLogReq, opts ...client.CallOption) (*FootLogRes, error) {
	req := c.c.NewRequest(c.name, "DatanodeService.FootLog", in)
	out := new(FootLogRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DatanodeService service

type DatanodeServiceHandler interface {
	// ---------- 用户接口 ------- //
	// 创建新用户
	CreateUser(context.Context, *NewUserReq, *UserRes) error
	// 获取用户信息
	UserInfo(context.Context, *UserInfoReq, *UserRes) error
	// 搜索用户信息
	SearchUser(context.Context, *SearchUserReq, *SearchUserRes) error
	// 更新用户信息
	UpdateUserInfo(context.Context, *UpdateUserinfoReq, *UpdateUserinfoRes) error
	// ---------------- 文件 ------------------//
	// 获取目录,目录之间的/替换成-
	FileDirOne(context.Context, *FileDirReq, *FileDirRes) error
	// 添加目录
	FileDirAdd(context.Context, *AddDirReq, *AddDirRes) error
	// 添加文件
	AddFile(context.Context, *AddFileReq, *AddFileRes) error
	// 获取云盘属性，容量和文件数量
	YunSaveAttr(context.Context, *FileAttrReq, *FileAttrRes) error
	// 获取图像缩略图
	Thumbnail(context.Context, *ThumbnailReq, *ThumbnailRes) error
	// ------------- key-value ----------------//
	// 设置缓存
	SetKeyValue(context.Context, *SetKeyvalueReq, *SetKeyvalueRes) error
	// 获取缓存
	GetKeyValue(context.Context, *GetKeyvalueReq, *GetKeyvalueRes) error
	// 删除
	DelKeyValue(context.Context, *DelKeyvalueReq, *DelKeyvalueRes) error
	//用户信息
	SetUserInfo(context.Context, *SetUserInfoReq, *SetUserInfoRes) error
	GetUserInfo(context.Context, *GetUserInfoReq, *GetUserInfoRes) error
	//微信用户信息操作
	SetWechatUinfo(context.Context, *SetWechatUinfoReq, *SetWechatUinfoRes) error
	GetWechatUinfo(context.Context, *GetWechatUinfoReq, *GetWechatUinfoRes) error
	// 支付宝用户信息操作
	SetAlipayUinfo(context.Context, *SetAlipayUinfoReq, *SetAlipayUinfoRes) error
	GetAlipayUinfo(context.Context, *GetAlipayUinfoReq, *GetAlipayUinfoRes) error
	// 请求百度实体信息
	SetBaiduEntity(context.Context, *SetBaiduEntityReq, *SetBaiduEntityRes) error
	BaiduEntity(context.Context, *BaiduEntityReq, *BaiduEntityRes) error
	// 提交实名
	SubmitCert(context.Context, *SubmitCertReq, *SubmitCertRes) error
	// 更新实名状态
	CertStatus(context.Context, *CertStatusReq, *CertStatusRes) error
	// 获取实名信息
	CertInfo(context.Context, *CertInfoReq, *CertInfoRes) error
	// 实名日志记录
	CertFlow(context.Context, *CertFlowReq, *CertFlowRes) error
	// 添加日志
	AddUserLog(context.Context, *AddUserLogReq, *AddUserLogRes) error
	// 获取日志
	QueryUserLog(context.Context, *QueryUserLogReq, *QueryUserLogRes) error
	// 日志数量
	UserLogCount(context.Context, *UserLogCountReq, *UserLogCountRes) error
	// 添加任务
	AddTask(context.Context, *AddTaskReq, *AddTaskRes) error
	// 获取任务
	GetTask(context.Context, *GetTaskReq, *GetTaskRes) error
	// 获取任务数量
	GetTaskCount(context.Context, *GetTaskCountReq, *GetTaskCountRes) error
	// 添加包信息
	AddPkg(context.Context, *AddPkgReq, *AddPkgRes) error
	// 删除包
	DelPkg(context.Context, *DelPkgReq, *DelPkgRes) error
	// 获取包列表
	PkgLists(context.Context, *PkgListReq, *PkgListRes) error
	// 获取包数量
	PkgCount(context.Context, *PkgCountReq, *PkgCountRes) error
	// 获取vip对应的限制
	VipLimit(context.Context, *VipLimitReq, *VipLimitRes) error
	// vip产品列表
	VipProduct(context.Context, *VipProductReq, *VipProductRes) error
	// vip产品介绍
	VipDesc(context.Context, *VipDescReq, *VipDescRes) error
	// vip充值
	VipPay(context.Context, *VipPayReq, *VipPayRes) error
	// 更新vip充值状态
	VipOrderStatus(context.Context, *VipOrderStatusReq, *VipOrderStatusRes) error
	// 获取vip充值列表
	VipOrderList(context.Context, *VipOrderListReq, *VipOrderListRes) error
	// 用户积分操作
	UserScore(context.Context, *UserScoreReq, *UserScoreRes) error
	// 用户vip操作
	UserVip(context.Context, *UserVipReq, *UserVipRes) error
	RefreshUserCoin(context.Context, *RefreshUserCoinReq, *RefreshUserCoinRes) error
	// 服务调用日志
	FootLog(context.Context, *FootLogReq, *FootLogRes) error
}

func RegisterDatanodeServiceHandler(s server.Server, hdlr DatanodeServiceHandler, opts ...server.HandlerOption) error {
	type datanodeService interface {
		CreateUser(ctx context.Context, in *NewUserReq, out *UserRes) error
		UserInfo(ctx context.Context, in *UserInfoReq, out *UserRes) error
		SearchUser(ctx context.Context, in *SearchUserReq, out *SearchUserRes) error
		UpdateUserInfo(ctx context.Context, in *UpdateUserinfoReq, out *UpdateUserinfoRes) error
		FileDirOne(ctx context.Context, in *FileDirReq, out *FileDirRes) error
		FileDirAdd(ctx context.Context, in *AddDirReq, out *AddDirRes) error
		AddFile(ctx context.Context, in *AddFileReq, out *AddFileRes) error
		YunSaveAttr(ctx context.Context, in *FileAttrReq, out *FileAttrRes) error
		Thumbnail(ctx context.Context, in *ThumbnailReq, out *ThumbnailRes) error
		SetKeyValue(ctx context.Context, in *SetKeyvalueReq, out *SetKeyvalueRes) error
		GetKeyValue(ctx context.Context, in *GetKeyvalueReq, out *GetKeyvalueRes) error
		DelKeyValue(ctx context.Context, in *DelKeyvalueReq, out *DelKeyvalueRes) error
		SetUserInfo(ctx context.Context, in *SetUserInfoReq, out *SetUserInfoRes) error
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, out *GetUserInfoRes) error
		SetWechatUinfo(ctx context.Context, in *SetWechatUinfoReq, out *SetWechatUinfoRes) error
		GetWechatUinfo(ctx context.Context, in *GetWechatUinfoReq, out *GetWechatUinfoRes) error
		SetAlipayUinfo(ctx context.Context, in *SetAlipayUinfoReq, out *SetAlipayUinfoRes) error
		GetAlipayUinfo(ctx context.Context, in *GetAlipayUinfoReq, out *GetAlipayUinfoRes) error
		SetBaiduEntity(ctx context.Context, in *SetBaiduEntityReq, out *SetBaiduEntityRes) error
		BaiduEntity(ctx context.Context, in *BaiduEntityReq, out *BaiduEntityRes) error
		SubmitCert(ctx context.Context, in *SubmitCertReq, out *SubmitCertRes) error
		CertStatus(ctx context.Context, in *CertStatusReq, out *CertStatusRes) error
		CertInfo(ctx context.Context, in *CertInfoReq, out *CertInfoRes) error
		CertFlow(ctx context.Context, in *CertFlowReq, out *CertFlowRes) error
		AddUserLog(ctx context.Context, in *AddUserLogReq, out *AddUserLogRes) error
		QueryUserLog(ctx context.Context, in *QueryUserLogReq, out *QueryUserLogRes) error
		UserLogCount(ctx context.Context, in *UserLogCountReq, out *UserLogCountRes) error
		AddTask(ctx context.Context, in *AddTaskReq, out *AddTaskRes) error
		GetTask(ctx context.Context, in *GetTaskReq, out *GetTaskRes) error
		GetTaskCount(ctx context.Context, in *GetTaskCountReq, out *GetTaskCountRes) error
		AddPkg(ctx context.Context, in *AddPkgReq, out *AddPkgRes) error
		DelPkg(ctx context.Context, in *DelPkgReq, out *DelPkgRes) error
		PkgLists(ctx context.Context, in *PkgListReq, out *PkgListRes) error
		PkgCount(ctx context.Context, in *PkgCountReq, out *PkgCountRes) error
		VipLimit(ctx context.Context, in *VipLimitReq, out *VipLimitRes) error
		VipProduct(ctx context.Context, in *VipProductReq, out *VipProductRes) error
		VipDesc(ctx context.Context, in *VipDescReq, out *VipDescRes) error
		VipPay(ctx context.Context, in *VipPayReq, out *VipPayRes) error
		VipOrderStatus(ctx context.Context, in *VipOrderStatusReq, out *VipOrderStatusRes) error
		VipOrderList(ctx context.Context, in *VipOrderListReq, out *VipOrderListRes) error
		UserScore(ctx context.Context, in *UserScoreReq, out *UserScoreRes) error
		UserVip(ctx context.Context, in *UserVipReq, out *UserVipRes) error
		RefreshUserCoin(ctx context.Context, in *RefreshUserCoinReq, out *RefreshUserCoinRes) error
		FootLog(ctx context.Context, in *FootLogReq, out *FootLogRes) error
	}
	type DatanodeService struct {
		datanodeService
	}
	h := &datanodeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DatanodeService{h}, opts...))
}

type datanodeServiceHandler struct {
	DatanodeServiceHandler
}

func (h *datanodeServiceHandler) CreateUser(ctx context.Context, in *NewUserReq, out *UserRes) error {
	return h.DatanodeServiceHandler.CreateUser(ctx, in, out)
}

func (h *datanodeServiceHandler) UserInfo(ctx context.Context, in *UserInfoReq, out *UserRes) error {
	return h.DatanodeServiceHandler.UserInfo(ctx, in, out)
}

func (h *datanodeServiceHandler) SearchUser(ctx context.Context, in *SearchUserReq, out *SearchUserRes) error {
	return h.DatanodeServiceHandler.SearchUser(ctx, in, out)
}

func (h *datanodeServiceHandler) UpdateUserInfo(ctx context.Context, in *UpdateUserinfoReq, out *UpdateUserinfoRes) error {
	return h.DatanodeServiceHandler.UpdateUserInfo(ctx, in, out)
}

func (h *datanodeServiceHandler) FileDirOne(ctx context.Context, in *FileDirReq, out *FileDirRes) error {
	return h.DatanodeServiceHandler.FileDirOne(ctx, in, out)
}

func (h *datanodeServiceHandler) FileDirAdd(ctx context.Context, in *AddDirReq, out *AddDirRes) error {
	return h.DatanodeServiceHandler.FileDirAdd(ctx, in, out)
}

func (h *datanodeServiceHandler) AddFile(ctx context.Context, in *AddFileReq, out *AddFileRes) error {
	return h.DatanodeServiceHandler.AddFile(ctx, in, out)
}

func (h *datanodeServiceHandler) YunSaveAttr(ctx context.Context, in *FileAttrReq, out *FileAttrRes) error {
	return h.DatanodeServiceHandler.YunSaveAttr(ctx, in, out)
}

func (h *datanodeServiceHandler) Thumbnail(ctx context.Context, in *ThumbnailReq, out *ThumbnailRes) error {
	return h.DatanodeServiceHandler.Thumbnail(ctx, in, out)
}

func (h *datanodeServiceHandler) SetKeyValue(ctx context.Context, in *SetKeyvalueReq, out *SetKeyvalueRes) error {
	return h.DatanodeServiceHandler.SetKeyValue(ctx, in, out)
}

func (h *datanodeServiceHandler) GetKeyValue(ctx context.Context, in *GetKeyvalueReq, out *GetKeyvalueRes) error {
	return h.DatanodeServiceHandler.GetKeyValue(ctx, in, out)
}

func (h *datanodeServiceHandler) DelKeyValue(ctx context.Context, in *DelKeyvalueReq, out *DelKeyvalueRes) error {
	return h.DatanodeServiceHandler.DelKeyValue(ctx, in, out)
}

func (h *datanodeServiceHandler) SetUserInfo(ctx context.Context, in *SetUserInfoReq, out *SetUserInfoRes) error {
	return h.DatanodeServiceHandler.SetUserInfo(ctx, in, out)
}

func (h *datanodeServiceHandler) GetUserInfo(ctx context.Context, in *GetUserInfoReq, out *GetUserInfoRes) error {
	return h.DatanodeServiceHandler.GetUserInfo(ctx, in, out)
}

func (h *datanodeServiceHandler) SetWechatUinfo(ctx context.Context, in *SetWechatUinfoReq, out *SetWechatUinfoRes) error {
	return h.DatanodeServiceHandler.SetWechatUinfo(ctx, in, out)
}

func (h *datanodeServiceHandler) GetWechatUinfo(ctx context.Context, in *GetWechatUinfoReq, out *GetWechatUinfoRes) error {
	return h.DatanodeServiceHandler.GetWechatUinfo(ctx, in, out)
}

func (h *datanodeServiceHandler) SetAlipayUinfo(ctx context.Context, in *SetAlipayUinfoReq, out *SetAlipayUinfoRes) error {
	return h.DatanodeServiceHandler.SetAlipayUinfo(ctx, in, out)
}

func (h *datanodeServiceHandler) GetAlipayUinfo(ctx context.Context, in *GetAlipayUinfoReq, out *GetAlipayUinfoRes) error {
	return h.DatanodeServiceHandler.GetAlipayUinfo(ctx, in, out)
}

func (h *datanodeServiceHandler) SetBaiduEntity(ctx context.Context, in *SetBaiduEntityReq, out *SetBaiduEntityRes) error {
	return h.DatanodeServiceHandler.SetBaiduEntity(ctx, in, out)
}

func (h *datanodeServiceHandler) BaiduEntity(ctx context.Context, in *BaiduEntityReq, out *BaiduEntityRes) error {
	return h.DatanodeServiceHandler.BaiduEntity(ctx, in, out)
}

func (h *datanodeServiceHandler) SubmitCert(ctx context.Context, in *SubmitCertReq, out *SubmitCertRes) error {
	return h.DatanodeServiceHandler.SubmitCert(ctx, in, out)
}

func (h *datanodeServiceHandler) CertStatus(ctx context.Context, in *CertStatusReq, out *CertStatusRes) error {
	return h.DatanodeServiceHandler.CertStatus(ctx, in, out)
}

func (h *datanodeServiceHandler) CertInfo(ctx context.Context, in *CertInfoReq, out *CertInfoRes) error {
	return h.DatanodeServiceHandler.CertInfo(ctx, in, out)
}

func (h *datanodeServiceHandler) CertFlow(ctx context.Context, in *CertFlowReq, out *CertFlowRes) error {
	return h.DatanodeServiceHandler.CertFlow(ctx, in, out)
}

func (h *datanodeServiceHandler) AddUserLog(ctx context.Context, in *AddUserLogReq, out *AddUserLogRes) error {
	return h.DatanodeServiceHandler.AddUserLog(ctx, in, out)
}

func (h *datanodeServiceHandler) QueryUserLog(ctx context.Context, in *QueryUserLogReq, out *QueryUserLogRes) error {
	return h.DatanodeServiceHandler.QueryUserLog(ctx, in, out)
}

func (h *datanodeServiceHandler) UserLogCount(ctx context.Context, in *UserLogCountReq, out *UserLogCountRes) error {
	return h.DatanodeServiceHandler.UserLogCount(ctx, in, out)
}

func (h *datanodeServiceHandler) AddTask(ctx context.Context, in *AddTaskReq, out *AddTaskRes) error {
	return h.DatanodeServiceHandler.AddTask(ctx, in, out)
}

func (h *datanodeServiceHandler) GetTask(ctx context.Context, in *GetTaskReq, out *GetTaskRes) error {
	return h.DatanodeServiceHandler.GetTask(ctx, in, out)
}

func (h *datanodeServiceHandler) GetTaskCount(ctx context.Context, in *GetTaskCountReq, out *GetTaskCountRes) error {
	return h.DatanodeServiceHandler.GetTaskCount(ctx, in, out)
}

func (h *datanodeServiceHandler) AddPkg(ctx context.Context, in *AddPkgReq, out *AddPkgRes) error {
	return h.DatanodeServiceHandler.AddPkg(ctx, in, out)
}

func (h *datanodeServiceHandler) DelPkg(ctx context.Context, in *DelPkgReq, out *DelPkgRes) error {
	return h.DatanodeServiceHandler.DelPkg(ctx, in, out)
}

func (h *datanodeServiceHandler) PkgLists(ctx context.Context, in *PkgListReq, out *PkgListRes) error {
	return h.DatanodeServiceHandler.PkgLists(ctx, in, out)
}

func (h *datanodeServiceHandler) PkgCount(ctx context.Context, in *PkgCountReq, out *PkgCountRes) error {
	return h.DatanodeServiceHandler.PkgCount(ctx, in, out)
}

func (h *datanodeServiceHandler) VipLimit(ctx context.Context, in *VipLimitReq, out *VipLimitRes) error {
	return h.DatanodeServiceHandler.VipLimit(ctx, in, out)
}

func (h *datanodeServiceHandler) VipProduct(ctx context.Context, in *VipProductReq, out *VipProductRes) error {
	return h.DatanodeServiceHandler.VipProduct(ctx, in, out)
}

func (h *datanodeServiceHandler) VipDesc(ctx context.Context, in *VipDescReq, out *VipDescRes) error {
	return h.DatanodeServiceHandler.VipDesc(ctx, in, out)
}

func (h *datanodeServiceHandler) VipPay(ctx context.Context, in *VipPayReq, out *VipPayRes) error {
	return h.DatanodeServiceHandler.VipPay(ctx, in, out)
}

func (h *datanodeServiceHandler) VipOrderStatus(ctx context.Context, in *VipOrderStatusReq, out *VipOrderStatusRes) error {
	return h.DatanodeServiceHandler.VipOrderStatus(ctx, in, out)
}

func (h *datanodeServiceHandler) VipOrderList(ctx context.Context, in *VipOrderListReq, out *VipOrderListRes) error {
	return h.DatanodeServiceHandler.VipOrderList(ctx, in, out)
}

func (h *datanodeServiceHandler) UserScore(ctx context.Context, in *UserScoreReq, out *UserScoreRes) error {
	return h.DatanodeServiceHandler.UserScore(ctx, in, out)
}

func (h *datanodeServiceHandler) UserVip(ctx context.Context, in *UserVipReq, out *UserVipRes) error {
	return h.DatanodeServiceHandler.UserVip(ctx, in, out)
}

func (h *datanodeServiceHandler) RefreshUserCoin(ctx context.Context, in *RefreshUserCoinReq, out *RefreshUserCoinRes) error {
	return h.DatanodeServiceHandler.RefreshUserCoin(ctx, in, out)
}

func (h *datanodeServiceHandler) FootLog(ctx context.Context, in *FootLogReq, out *FootLogRes) error {
	return h.DatanodeServiceHandler.FootLog(ctx, in, out)
}
