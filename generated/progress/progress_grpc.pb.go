// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
// source: progress/progress.proto

package progress

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProgressService_GetUserProgress_FullMethodName    = "/progress.ProgressService/GetUserProgress"
	ProgressService_GetDailyProgress_FullMethodName   = "/progress.ProgressService/GetDailyProgress"
	ProgressService_GetWeeklyProgress_FullMethodName  = "/progress.ProgressService/GetWeeklyProgress"
	ProgressService_GetMonthlyProgress_FullMethodName = "/progress.ProgressService/GetMonthlyProgress"
	ProgressService_GetUserAchievement_FullMethodName = "/progress.ProgressService/GetUserAchievement"
	ProgressService_GetLeaders_FullMethodName         = "/progress.ProgressService/GetLeaders"
	ProgressService_GetUserSkills_FullMethodName      = "/progress.ProgressService/GetUserSkills"
	ProgressService_GetStatistics_FullMethodName      = "/progress.ProgressService/GetStatistics"
)

// ProgressServiceClient is the client API for ProgressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgressServiceClient interface {
	GetUserProgress(ctx context.Context, in *LCodeUID, opts ...grpc.CallOption) (*GetUserProgressResponse, error)
	GetDailyProgress(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*GetDailyProgressResponse, error)
	GetWeeklyProgress(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*GetWeeklyProgressResponse, error)
	GetMonthlyProgress(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*GetMonthlyProgressResponse, error)
	GetUserAchievement(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*GetUserAchievementsResponse, error)
	GetLeaders(ctx context.Context, in *LanguageCode, opts ...grpc.CallOption) (*GetLeadersResponse, error)
	GetUserSkills(ctx context.Context, in *LCodeUID, opts ...grpc.CallOption) (*GetUserSkillsResponse, error)
	GetStatistics(ctx context.Context, in *LCodeUID, opts ...grpc.CallOption) (*GetStatisticsResponse, error)
}

type progressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProgressServiceClient(cc grpc.ClientConnInterface) ProgressServiceClient {
	return &progressServiceClient{cc}
}

func (c *progressServiceClient) GetUserProgress(ctx context.Context, in *LCodeUID, opts ...grpc.CallOption) (*GetUserProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetUserProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetDailyProgress(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*GetDailyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDailyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetDailyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetWeeklyProgress(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*GetWeeklyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWeeklyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetWeeklyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetMonthlyProgress(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*GetMonthlyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMonthlyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetMonthlyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetUserAchievement(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*GetUserAchievementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserAchievementsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetUserAchievement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetLeaders(ctx context.Context, in *LanguageCode, opts ...grpc.CallOption) (*GetLeadersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLeadersResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetLeaders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetUserSkills(ctx context.Context, in *LCodeUID, opts ...grpc.CallOption) (*GetUserSkillsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserSkillsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetUserSkills_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetStatistics(ctx context.Context, in *LCodeUID, opts ...grpc.CallOption) (*GetStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatisticsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProgressServiceServer is the server API for ProgressService service.
// All implementations must embed UnimplementedProgressServiceServer
// for forward compatibility
type ProgressServiceServer interface {
	GetUserProgress(context.Context, *LCodeUID) (*GetUserProgressResponse, error)
	GetDailyProgress(context.Context, *UserID) (*GetDailyProgressResponse, error)
	GetWeeklyProgress(context.Context, *UserID) (*GetWeeklyProgressResponse, error)
	GetMonthlyProgress(context.Context, *UserID) (*GetMonthlyProgressResponse, error)
	GetUserAchievement(context.Context, *UserID) (*GetUserAchievementsResponse, error)
	GetLeaders(context.Context, *LanguageCode) (*GetLeadersResponse, error)
	GetUserSkills(context.Context, *LCodeUID) (*GetUserSkillsResponse, error)
	GetStatistics(context.Context, *LCodeUID) (*GetStatisticsResponse, error)
	mustEmbedUnimplementedProgressServiceServer()
}

// UnimplementedProgressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProgressServiceServer struct {
}

func (UnimplementedProgressServiceServer) GetUserProgress(context.Context, *LCodeUID) (*GetUserProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetDailyProgress(context.Context, *UserID) (*GetDailyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDailyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetWeeklyProgress(context.Context, *UserID) (*GetWeeklyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeeklyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetMonthlyProgress(context.Context, *UserID) (*GetMonthlyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetUserAchievement(context.Context, *UserID) (*GetUserAchievementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAchievement not implemented")
}
func (UnimplementedProgressServiceServer) GetLeaders(context.Context, *LanguageCode) (*GetLeadersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaders not implemented")
}
func (UnimplementedProgressServiceServer) GetUserSkills(context.Context, *LCodeUID) (*GetUserSkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSkills not implemented")
}
func (UnimplementedProgressServiceServer) GetStatistics(context.Context, *LCodeUID) (*GetStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedProgressServiceServer) mustEmbedUnimplementedProgressServiceServer() {}

// UnsafeProgressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgressServiceServer will
// result in compilation errors.
type UnsafeProgressServiceServer interface {
	mustEmbedUnimplementedProgressServiceServer()
}

func RegisterProgressServiceServer(s grpc.ServiceRegistrar, srv ProgressServiceServer) {
	s.RegisterService(&ProgressService_ServiceDesc, srv)
}

func _ProgressService_GetUserProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LCodeUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetUserProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetUserProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetUserProgress(ctx, req.(*LCodeUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetDailyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetDailyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetDailyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetDailyProgress(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetWeeklyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetWeeklyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetWeeklyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetWeeklyProgress(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetMonthlyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetMonthlyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetMonthlyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetMonthlyProgress(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetUserAchievement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetUserAchievement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetUserAchievement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetUserAchievement(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetLeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LanguageCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetLeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetLeaders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetLeaders(ctx, req.(*LanguageCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetUserSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LCodeUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetUserSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetUserSkills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetUserSkills(ctx, req.(*LCodeUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LCodeUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetStatistics(ctx, req.(*LCodeUID))
	}
	return interceptor(ctx, in, info, handler)
}

// ProgressService_ServiceDesc is the grpc.ServiceDesc for ProgressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "progress.ProgressService",
	HandlerType: (*ProgressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserProgress",
			Handler:    _ProgressService_GetUserProgress_Handler,
		},
		{
			MethodName: "GetDailyProgress",
			Handler:    _ProgressService_GetDailyProgress_Handler,
		},
		{
			MethodName: "GetWeeklyProgress",
			Handler:    _ProgressService_GetWeeklyProgress_Handler,
		},
		{
			MethodName: "GetMonthlyProgress",
			Handler:    _ProgressService_GetMonthlyProgress_Handler,
		},
		{
			MethodName: "GetUserAchievement",
			Handler:    _ProgressService_GetUserAchievement_Handler,
		},
		{
			MethodName: "GetLeaders",
			Handler:    _ProgressService_GetLeaders_Handler,
		},
		{
			MethodName: "GetUserSkills",
			Handler:    _ProgressService_GetUserSkills_Handler,
		},
		{
			MethodName: "GetStatistics",
			Handler:    _ProgressService_GetStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "progress/progress.proto",
}
