// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: foundation_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FoundationServiceClient is the client API for FoundationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoundationServiceClient interface {
	// * 检测是否包含敏感词
	CheckSensitive(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error)
	// * 替换敏感词
	ReplaceSensitive(ctx context.Context, in *ReplaceSensitiveRequest, opts ...grpc.CallOption) (*String, error)
	// * 清除缓存
	CleanCache(ctx context.Context, in *CleanCacheRequest, opts ...grpc.CallOption) (*CleanCacheResponse, error)
	// * 获取短信API凭据, provider 短信服务商, 默认:http
	GetSmsApi(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSmsApi, error)
	// * 保存短信API凭据,@provider 短信服务商, 默认:http
	SaveSmsApi(ctx context.Context, in *SmsApiSaveRequest, opts ...grpc.CallOption) (*Result, error)
	// * 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
	SaveBoardHook(ctx context.Context, in *BoardHookSaveRequest, opts ...grpc.CallOption) (*Result, error)
	// 格式化资源地址并返回
	ResourceUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 设置键值
	// rpc SetValue (Pair) returns (Result){}
	// 删除值,key
	// rpc DeleteValue (String) returns (Result){}
	// 根据前缀获取值,prefix
	// rpc GetValuesByPrefix (String) returns (StringMap){}
	// 注册单点登录应用,返回值：
	//   - 1. 成功，并返回token
	//   - -1. 接口地址不正确
	//   - -2. 已经注册
	RegisterApp(ctx context.Context, in *SSsoApp, opts ...grpc.CallOption) (*String, error)
	// 获取应用信息,name
	GetApp(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSsoApp, error)
	// 获取单点登录应用
	GetAllSsoApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListResponse, error)
	// 验证超级用户账号和密码
	SuperValidate(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*SuperLoginResponse, error)
	// 保存超级用户账号和密码
	FlushSuperPwd(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Result, error)
	// 创建同步登录的地址,returnUrl
	GetSyncLoginUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 获取地区名称
	GetAreaNames(ctx context.Context, in *GetAreaNamesRequest, opts ...grpc.CallOption) (*StringListResponse, error)
	// 获取省市区字符串
	GetAreaString(ctx context.Context, in *AreaStringRequest, opts ...grpc.CallOption) (*String, error)
	// 获取下级区域,code
	GetChildAreas(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*AreaListResponse, error)
	// 获取移动应用设置
	GetMoAppConf(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SMobileAppConfig, error)
	// 保存移动应用设置
	SaveMoAppConf(ctx context.Context, in *SMobileAppConfig, opts ...grpc.CallOption) (*Result, error)
	// 获取微信接口配置
	GetWxApiConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SWxApiConfig, error)
	// 保存微信接口配置
	SaveWxApiConfig(ctx context.Context, in *SWxApiConfig, opts ...grpc.CallOption) (*Result, error)
	// 获取支付平台
	GetPayPlatform(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PaymentPlatformResponse, error)
	// 获取全局商户销售设置
	GetGlobMchSaleConf_(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SGlobMchSaleConf, error)
	// 保存全局商户销售设置
	SaveGlobMchSaleConf_(ctx context.Context, in *SGlobMchSaleConf, opts ...grpc.CallOption) (*Result, error)
}

type foundationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFoundationServiceClient(cc grpc.ClientConnInterface) FoundationServiceClient {
	return &foundationServiceClient{cc}
}

func (c *foundationServiceClient) CheckSensitive(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/FoundationService/CheckSensitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) ReplaceSensitive(ctx context.Context, in *ReplaceSensitiveRequest, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/ReplaceSensitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) CleanCache(ctx context.Context, in *CleanCacheRequest, opts ...grpc.CallOption) (*CleanCacheResponse, error) {
	out := new(CleanCacheResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/CleanCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetSmsApi(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSmsApi, error) {
	out := new(SSmsApi)
	err := c.cc.Invoke(ctx, "/FoundationService/GetSmsApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveSmsApi(ctx context.Context, in *SmsApiSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveSmsApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveBoardHook(ctx context.Context, in *BoardHookSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveBoardHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) ResourceUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/ResourceUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) RegisterApp(ctx context.Context, in *SSsoApp, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/RegisterApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetApp(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSsoApp, error) {
	out := new(SSsoApp)
	err := c.cc.Invoke(ctx, "/FoundationService/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAllSsoApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListResponse, error) {
	out := new(StringListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAllSsoApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SuperValidate(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*SuperLoginResponse, error) {
	out := new(SuperLoginResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/SuperValidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) FlushSuperPwd(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/FlushSuperPwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetSyncLoginUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/GetSyncLoginUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAreaNames(ctx context.Context, in *GetAreaNamesRequest, opts ...grpc.CallOption) (*StringListResponse, error) {
	out := new(StringListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAreaNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAreaString(ctx context.Context, in *AreaStringRequest, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAreaString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetChildAreas(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*AreaListResponse, error) {
	out := new(AreaListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetChildAreas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetMoAppConf(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SMobileAppConfig, error) {
	out := new(SMobileAppConfig)
	err := c.cc.Invoke(ctx, "/FoundationService/GetMoAppConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveMoAppConf(ctx context.Context, in *SMobileAppConfig, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveMoAppConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetWxApiConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SWxApiConfig, error) {
	out := new(SWxApiConfig)
	err := c.cc.Invoke(ctx, "/FoundationService/GetWxApiConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveWxApiConfig(ctx context.Context, in *SWxApiConfig, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveWxApiConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetPayPlatform(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PaymentPlatformResponse, error) {
	out := new(PaymentPlatformResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetPayPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetGlobMchSaleConf_(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SGlobMchSaleConf, error) {
	out := new(SGlobMchSaleConf)
	err := c.cc.Invoke(ctx, "/FoundationService/GetGlobMchSaleConf_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveGlobMchSaleConf_(ctx context.Context, in *SGlobMchSaleConf, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveGlobMchSaleConf_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoundationServiceServer is the server API for FoundationService service.
// All implementations must embed UnimplementedFoundationServiceServer
// for forward compatibility
type FoundationServiceServer interface {
	// * 检测是否包含敏感词
	CheckSensitive(context.Context, *String) (*Bool, error)
	// * 替换敏感词
	ReplaceSensitive(context.Context, *ReplaceSensitiveRequest) (*String, error)
	// * 清除缓存
	CleanCache(context.Context, *CleanCacheRequest) (*CleanCacheResponse, error)
	// * 获取短信API凭据, provider 短信服务商, 默认:http
	GetSmsApi(context.Context, *String) (*SSmsApi, error)
	// * 保存短信API凭据,@provider 短信服务商, 默认:http
	SaveSmsApi(context.Context, *SmsApiSaveRequest) (*Result, error)
	// * 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
	SaveBoardHook(context.Context, *BoardHookSaveRequest) (*Result, error)
	// 格式化资源地址并返回
	ResourceUrl(context.Context, *String) (*String, error)
	// 设置键值
	// rpc SetValue (Pair) returns (Result){}
	// 删除值,key
	// rpc DeleteValue (String) returns (Result){}
	// 根据前缀获取值,prefix
	// rpc GetValuesByPrefix (String) returns (StringMap){}
	// 注册单点登录应用,返回值：
	//   - 1. 成功，并返回token
	//   - -1. 接口地址不正确
	//   - -2. 已经注册
	RegisterApp(context.Context, *SSsoApp) (*String, error)
	// 获取应用信息,name
	GetApp(context.Context, *String) (*SSsoApp, error)
	// 获取单点登录应用
	GetAllSsoApp(context.Context, *Empty) (*StringListResponse, error)
	// 验证超级用户账号和密码
	SuperValidate(context.Context, *UserPwd) (*SuperLoginResponse, error)
	// 保存超级用户账号和密码
	FlushSuperPwd(context.Context, *UserPwd) (*Result, error)
	// 创建同步登录的地址,returnUrl
	GetSyncLoginUrl(context.Context, *String) (*String, error)
	// 获取地区名称
	GetAreaNames(context.Context, *GetAreaNamesRequest) (*StringListResponse, error)
	// 获取省市区字符串
	GetAreaString(context.Context, *AreaStringRequest) (*String, error)
	// 获取下级区域,code
	GetChildAreas(context.Context, *Int32) (*AreaListResponse, error)
	// 获取移动应用设置
	GetMoAppConf(context.Context, *Empty) (*SMobileAppConfig, error)
	// 保存移动应用设置
	SaveMoAppConf(context.Context, *SMobileAppConfig) (*Result, error)
	// 获取微信接口配置
	GetWxApiConfig(context.Context, *Empty) (*SWxApiConfig, error)
	// 保存微信接口配置
	SaveWxApiConfig(context.Context, *SWxApiConfig) (*Result, error)
	// 获取支付平台
	GetPayPlatform(context.Context, *Empty) (*PaymentPlatformResponse, error)
	// 获取全局商户销售设置
	GetGlobMchSaleConf_(context.Context, *Empty) (*SGlobMchSaleConf, error)
	// 保存全局商户销售设置
	SaveGlobMchSaleConf_(context.Context, *SGlobMchSaleConf) (*Result, error)
	mustEmbedUnimplementedFoundationServiceServer()
}

// UnimplementedFoundationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFoundationServiceServer struct {
}

func (UnimplementedFoundationServiceServer) CheckSensitive(context.Context, *String) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSensitive not implemented")
}
func (UnimplementedFoundationServiceServer) ReplaceSensitive(context.Context, *ReplaceSensitiveRequest) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceSensitive not implemented")
}
func (UnimplementedFoundationServiceServer) CleanCache(context.Context, *CleanCacheRequest) (*CleanCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanCache not implemented")
}
func (UnimplementedFoundationServiceServer) GetSmsApi(context.Context, *String) (*SSmsApi, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmsApi not implemented")
}
func (UnimplementedFoundationServiceServer) SaveSmsApi(context.Context, *SmsApiSaveRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSmsApi not implemented")
}
func (UnimplementedFoundationServiceServer) SaveBoardHook(context.Context, *BoardHookSaveRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBoardHook not implemented")
}
func (UnimplementedFoundationServiceServer) ResourceUrl(context.Context, *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceUrl not implemented")
}
func (UnimplementedFoundationServiceServer) RegisterApp(context.Context, *SSsoApp) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterApp not implemented")
}
func (UnimplementedFoundationServiceServer) GetApp(context.Context, *String) (*SSsoApp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}
func (UnimplementedFoundationServiceServer) GetAllSsoApp(context.Context, *Empty) (*StringListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSsoApp not implemented")
}
func (UnimplementedFoundationServiceServer) SuperValidate(context.Context, *UserPwd) (*SuperLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuperValidate not implemented")
}
func (UnimplementedFoundationServiceServer) FlushSuperPwd(context.Context, *UserPwd) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushSuperPwd not implemented")
}
func (UnimplementedFoundationServiceServer) GetSyncLoginUrl(context.Context, *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSyncLoginUrl not implemented")
}
func (UnimplementedFoundationServiceServer) GetAreaNames(context.Context, *GetAreaNamesRequest) (*StringListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAreaNames not implemented")
}
func (UnimplementedFoundationServiceServer) GetAreaString(context.Context, *AreaStringRequest) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAreaString not implemented")
}
func (UnimplementedFoundationServiceServer) GetChildAreas(context.Context, *Int32) (*AreaListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildAreas not implemented")
}
func (UnimplementedFoundationServiceServer) GetMoAppConf(context.Context, *Empty) (*SMobileAppConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoAppConf not implemented")
}
func (UnimplementedFoundationServiceServer) SaveMoAppConf(context.Context, *SMobileAppConfig) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMoAppConf not implemented")
}
func (UnimplementedFoundationServiceServer) GetWxApiConfig(context.Context, *Empty) (*SWxApiConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWxApiConfig not implemented")
}
func (UnimplementedFoundationServiceServer) SaveWxApiConfig(context.Context, *SWxApiConfig) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveWxApiConfig not implemented")
}
func (UnimplementedFoundationServiceServer) GetPayPlatform(context.Context, *Empty) (*PaymentPlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayPlatform not implemented")
}
func (UnimplementedFoundationServiceServer) GetGlobMchSaleConf_(context.Context, *Empty) (*SGlobMchSaleConf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobMchSaleConf_ not implemented")
}
func (UnimplementedFoundationServiceServer) SaveGlobMchSaleConf_(context.Context, *SGlobMchSaleConf) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveGlobMchSaleConf_ not implemented")
}
func (UnimplementedFoundationServiceServer) mustEmbedUnimplementedFoundationServiceServer() {}

// UnsafeFoundationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoundationServiceServer will
// result in compilation errors.
type UnsafeFoundationServiceServer interface {
	mustEmbedUnimplementedFoundationServiceServer()
}

func RegisterFoundationServiceServer(s grpc.ServiceRegistrar, srv FoundationServiceServer) {
	s.RegisterService(&FoundationService_ServiceDesc, srv)
}

func _FoundationService_CheckSensitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).CheckSensitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/CheckSensitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).CheckSensitive(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_ReplaceSensitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceSensitiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).ReplaceSensitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/ReplaceSensitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).ReplaceSensitive(ctx, req.(*ReplaceSensitiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_CleanCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).CleanCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/CleanCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).CleanCache(ctx, req.(*CleanCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetSmsApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetSmsApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetSmsApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetSmsApi(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveSmsApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsApiSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveSmsApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveSmsApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveSmsApi(ctx, req.(*SmsApiSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveBoardHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardHookSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveBoardHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveBoardHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveBoardHook(ctx, req.(*BoardHookSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_ResourceUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).ResourceUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/ResourceUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).ResourceUrl(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_RegisterApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSsoApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).RegisterApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/RegisterApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).RegisterApp(ctx, req.(*SSsoApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetApp(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAllSsoApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAllSsoApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAllSsoApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAllSsoApp(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SuperValidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SuperValidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SuperValidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SuperValidate(ctx, req.(*UserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_FlushSuperPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).FlushSuperPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/FlushSuperPwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).FlushSuperPwd(ctx, req.(*UserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetSyncLoginUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetSyncLoginUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetSyncLoginUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetSyncLoginUrl(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAreaNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAreaNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAreaNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAreaNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAreaNames(ctx, req.(*GetAreaNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAreaString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAreaString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAreaString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAreaString(ctx, req.(*AreaStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetChildAreas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetChildAreas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetChildAreas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetChildAreas(ctx, req.(*Int32))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetMoAppConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetMoAppConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetMoAppConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetMoAppConf(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveMoAppConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMobileAppConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveMoAppConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveMoAppConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveMoAppConf(ctx, req.(*SMobileAppConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetWxApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetWxApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetWxApiConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetWxApiConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveWxApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SWxApiConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveWxApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveWxApiConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveWxApiConfig(ctx, req.(*SWxApiConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetPayPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetPayPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetPayPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetPayPlatform(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetGlobMchSaleConf__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetGlobMchSaleConf_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetGlobMchSaleConf_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetGlobMchSaleConf_(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveGlobMchSaleConf__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SGlobMchSaleConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveGlobMchSaleConf_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveGlobMchSaleConf_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveGlobMchSaleConf_(ctx, req.(*SGlobMchSaleConf))
	}
	return interceptor(ctx, in, info, handler)
}

// FoundationService_ServiceDesc is the grpc.ServiceDesc for FoundationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FoundationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FoundationService",
	HandlerType: (*FoundationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckSensitive",
			Handler:    _FoundationService_CheckSensitive_Handler,
		},
		{
			MethodName: "ReplaceSensitive",
			Handler:    _FoundationService_ReplaceSensitive_Handler,
		},
		{
			MethodName: "CleanCache",
			Handler:    _FoundationService_CleanCache_Handler,
		},
		{
			MethodName: "GetSmsApi",
			Handler:    _FoundationService_GetSmsApi_Handler,
		},
		{
			MethodName: "SaveSmsApi",
			Handler:    _FoundationService_SaveSmsApi_Handler,
		},
		{
			MethodName: "SaveBoardHook",
			Handler:    _FoundationService_SaveBoardHook_Handler,
		},
		{
			MethodName: "ResourceUrl",
			Handler:    _FoundationService_ResourceUrl_Handler,
		},
		{
			MethodName: "RegisterApp",
			Handler:    _FoundationService_RegisterApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _FoundationService_GetApp_Handler,
		},
		{
			MethodName: "GetAllSsoApp",
			Handler:    _FoundationService_GetAllSsoApp_Handler,
		},
		{
			MethodName: "SuperValidate",
			Handler:    _FoundationService_SuperValidate_Handler,
		},
		{
			MethodName: "FlushSuperPwd",
			Handler:    _FoundationService_FlushSuperPwd_Handler,
		},
		{
			MethodName: "GetSyncLoginUrl",
			Handler:    _FoundationService_GetSyncLoginUrl_Handler,
		},
		{
			MethodName: "GetAreaNames",
			Handler:    _FoundationService_GetAreaNames_Handler,
		},
		{
			MethodName: "GetAreaString",
			Handler:    _FoundationService_GetAreaString_Handler,
		},
		{
			MethodName: "GetChildAreas",
			Handler:    _FoundationService_GetChildAreas_Handler,
		},
		{
			MethodName: "GetMoAppConf",
			Handler:    _FoundationService_GetMoAppConf_Handler,
		},
		{
			MethodName: "SaveMoAppConf",
			Handler:    _FoundationService_SaveMoAppConf_Handler,
		},
		{
			MethodName: "GetWxApiConfig",
			Handler:    _FoundationService_GetWxApiConfig_Handler,
		},
		{
			MethodName: "SaveWxApiConfig",
			Handler:    _FoundationService_SaveWxApiConfig_Handler,
		},
		{
			MethodName: "GetPayPlatform",
			Handler:    _FoundationService_GetPayPlatform_Handler,
		},
		{
			MethodName: "GetGlobMchSaleConf_",
			Handler:    _FoundationService_GetGlobMchSaleConf__Handler,
		},
		{
			MethodName: "SaveGlobMchSaleConf_",
			Handler:    _FoundationService_SaveGlobMchSaleConf__Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foundation_service.proto",
}
