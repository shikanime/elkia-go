// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/api/fleet/v1alpha1/fleet.proto

package v1alpha1

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

const (
	Cluster_MemberAdd_FullMethodName    = "/shikanime.elkia.fleet.v1alpha1.Cluster/MemberAdd"
	Cluster_MemberRemove_FullMethodName = "/shikanime.elkia.fleet.v1alpha1.Cluster/MemberRemove"
	Cluster_MemberUpdate_FullMethodName = "/shikanime.elkia.fleet.v1alpha1.Cluster/MemberUpdate"
	Cluster_MemberList_FullMethodName   = "/shikanime.elkia.fleet.v1alpha1.Cluster/MemberList"
)

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterClient interface {
	// MemberAdd adds a member into the cluster.
	MemberAdd(ctx context.Context, in *MemberAddRequest, opts ...grpc.CallOption) (*MemberAddResponse, error)
	// MemberRemove removes an existing member from the cluster.
	MemberRemove(ctx context.Context, in *MemberRemoveRequest, opts ...grpc.CallOption) (*MemberRemoveResponse, error)
	// MemberUpdate updates the member configuration.
	MemberUpdate(ctx context.Context, in *MemberUpdateRequest, opts ...grpc.CallOption) (*MemberUpdateResponse, error)
	// MemberList lists all the members in the cluster.
	MemberList(ctx context.Context, in *MemberListRequest, opts ...grpc.CallOption) (*MemberListResponse, error)
}

type clusterClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterClient(cc grpc.ClientConnInterface) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) MemberAdd(ctx context.Context, in *MemberAddRequest, opts ...grpc.CallOption) (*MemberAddResponse, error) {
	out := new(MemberAddResponse)
	err := c.cc.Invoke(ctx, Cluster_MemberAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) MemberRemove(ctx context.Context, in *MemberRemoveRequest, opts ...grpc.CallOption) (*MemberRemoveResponse, error) {
	out := new(MemberRemoveResponse)
	err := c.cc.Invoke(ctx, Cluster_MemberRemove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) MemberUpdate(ctx context.Context, in *MemberUpdateRequest, opts ...grpc.CallOption) (*MemberUpdateResponse, error) {
	out := new(MemberUpdateResponse)
	err := c.cc.Invoke(ctx, Cluster_MemberUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) MemberList(ctx context.Context, in *MemberListRequest, opts ...grpc.CallOption) (*MemberListResponse, error) {
	out := new(MemberListResponse)
	err := c.cc.Invoke(ctx, Cluster_MemberList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
// All implementations must embed UnimplementedClusterServer
// for forward compatibility
type ClusterServer interface {
	// MemberAdd adds a member into the cluster.
	MemberAdd(context.Context, *MemberAddRequest) (*MemberAddResponse, error)
	// MemberRemove removes an existing member from the cluster.
	MemberRemove(context.Context, *MemberRemoveRequest) (*MemberRemoveResponse, error)
	// MemberUpdate updates the member configuration.
	MemberUpdate(context.Context, *MemberUpdateRequest) (*MemberUpdateResponse, error)
	// MemberList lists all the members in the cluster.
	MemberList(context.Context, *MemberListRequest) (*MemberListResponse, error)
	mustEmbedUnimplementedClusterServer()
}

// UnimplementedClusterServer must be embedded to have forward compatible implementations.
type UnimplementedClusterServer struct {
}

func (UnimplementedClusterServer) MemberAdd(context.Context, *MemberAddRequest) (*MemberAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberAdd not implemented")
}
func (UnimplementedClusterServer) MemberRemove(context.Context, *MemberRemoveRequest) (*MemberRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberRemove not implemented")
}
func (UnimplementedClusterServer) MemberUpdate(context.Context, *MemberUpdateRequest) (*MemberUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberUpdate not implemented")
}
func (UnimplementedClusterServer) MemberList(context.Context, *MemberListRequest) (*MemberListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberList not implemented")
}
func (UnimplementedClusterServer) mustEmbedUnimplementedClusterServer() {}

// UnsafeClusterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServer will
// result in compilation errors.
type UnsafeClusterServer interface {
	mustEmbedUnimplementedClusterServer()
}

func RegisterClusterServer(s grpc.ServiceRegistrar, srv ClusterServer) {
	s.RegisterService(&Cluster_ServiceDesc, srv)
}

func _Cluster_MemberAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cluster_MemberAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberAdd(ctx, req.(*MemberAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_MemberRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cluster_MemberRemove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberRemove(ctx, req.(*MemberRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_MemberUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cluster_MemberUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberUpdate(ctx, req.(*MemberUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_MemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cluster_MemberList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberList(ctx, req.(*MemberListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cluster_ServiceDesc is the grpc.ServiceDesc for Cluster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cluster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shikanime.elkia.fleet.v1alpha1.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MemberAdd",
			Handler:    _Cluster_MemberAdd_Handler,
		},
		{
			MethodName: "MemberRemove",
			Handler:    _Cluster_MemberRemove_Handler,
		},
		{
			MethodName: "MemberUpdate",
			Handler:    _Cluster_MemberUpdate_Handler,
		},
		{
			MethodName: "MemberList",
			Handler:    _Cluster_MemberList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/fleet/v1alpha1/fleet.proto",
}

const (
	Presence_AuthCreateHandoffFlow_FullMethodName   = "/shikanime.elkia.fleet.v1alpha1.Presence/AuthCreateHandoffFlow"
	Presence_AuthCompleteHandoffFlow_FullMethodName = "/shikanime.elkia.fleet.v1alpha1.Presence/AuthCompleteHandoffFlow"
	Presence_AuthLogin_FullMethodName               = "/shikanime.elkia.fleet.v1alpha1.Presence/AuthLogin"
	Presence_AuthRefreshLogin_FullMethodName        = "/shikanime.elkia.fleet.v1alpha1.Presence/AuthRefreshLogin"
	Presence_AuthWhoAmI_FullMethodName              = "/shikanime.elkia.fleet.v1alpha1.Presence/AuthWhoAmI"
	Presence_AuthLogout_FullMethodName              = "/shikanime.elkia.fleet.v1alpha1.Presence/AuthLogout"
	Presence_SessionGet_FullMethodName              = "/shikanime.elkia.fleet.v1alpha1.Presence/SessionGet"
	Presence_SessionPut_FullMethodName              = "/shikanime.elkia.fleet.v1alpha1.Presence/SessionPut"
	Presence_SessionDelete_FullMethodName           = "/shikanime.elkia.fleet.v1alpha1.Presence/SessionDelete"
)

// PresenceClient is the client API for Presence service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PresenceClient interface {
	// AuthCreateHandoffFlow creates a handoff flow with a given identifier and
	// password.
	AuthCreateHandoffFlow(ctx context.Context, in *AuthCreateHandoffFlowRequest, opts ...grpc.CallOption) (*AuthCreateHandoffFlowResponse, error)
	// AuthCompleteHandoffFlow hands off a session to a gateway with a given token and
	// code.
	AuthCompleteHandoffFlow(ctx context.Context, in *AuthCompleteHandoffFlowRequest, opts ...grpc.CallOption) (*AuthCompleteHandoffFlowResponse, error)
	// AuthLogin hands off a session to a gateway with a given token and
	// code.
	AuthLogin(ctx context.Context, in *AuthLoginRequest, opts ...grpc.CallOption) (*AuthLoginResponse, error)
	// AuthRefreshLogin authenticates a gateway with a given identifier, password, and
	// token.
	AuthRefreshLogin(ctx context.Context, in *AuthRefreshLoginRequest, opts ...grpc.CallOption) (*AuthRefreshLoginResponse, error)
	// AuthWhoAmI returns the session associated with a given token.
	AuthWhoAmI(ctx context.Context, in *AuthWhoAmIRequest, opts ...grpc.CallOption) (*AuthWhoAmIResponse, error)
	// AuthLogout logs out a session with a given code.
	AuthLogout(ctx context.Context, in *AuthLogoutRequest, opts ...grpc.CallOption) (*AuthLogoutResponse, error)
	// SessionGet gets a session with a given code.
	SessionGet(ctx context.Context, in *SessionGetRequest, opts ...grpc.CallOption) (*SessionGetResponse, error)
	// SessionPut creates a session with a given identifier and token.
	SessionPut(ctx context.Context, in *SessionPutRequest, opts ...grpc.CallOption) (*SessionPutResponse, error)
	// SessionDelete Deletes a session with a given code.
	SessionDelete(ctx context.Context, in *SessionDeleteRequest, opts ...grpc.CallOption) (*SessionDeleteResponse, error)
}

type presenceClient struct {
	cc grpc.ClientConnInterface
}

func NewPresenceClient(cc grpc.ClientConnInterface) PresenceClient {
	return &presenceClient{cc}
}

func (c *presenceClient) AuthCreateHandoffFlow(ctx context.Context, in *AuthCreateHandoffFlowRequest, opts ...grpc.CallOption) (*AuthCreateHandoffFlowResponse, error) {
	out := new(AuthCreateHandoffFlowResponse)
	err := c.cc.Invoke(ctx, Presence_AuthCreateHandoffFlow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceClient) AuthCompleteHandoffFlow(ctx context.Context, in *AuthCompleteHandoffFlowRequest, opts ...grpc.CallOption) (*AuthCompleteHandoffFlowResponse, error) {
	out := new(AuthCompleteHandoffFlowResponse)
	err := c.cc.Invoke(ctx, Presence_AuthCompleteHandoffFlow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceClient) AuthLogin(ctx context.Context, in *AuthLoginRequest, opts ...grpc.CallOption) (*AuthLoginResponse, error) {
	out := new(AuthLoginResponse)
	err := c.cc.Invoke(ctx, Presence_AuthLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceClient) AuthRefreshLogin(ctx context.Context, in *AuthRefreshLoginRequest, opts ...grpc.CallOption) (*AuthRefreshLoginResponse, error) {
	out := new(AuthRefreshLoginResponse)
	err := c.cc.Invoke(ctx, Presence_AuthRefreshLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceClient) AuthWhoAmI(ctx context.Context, in *AuthWhoAmIRequest, opts ...grpc.CallOption) (*AuthWhoAmIResponse, error) {
	out := new(AuthWhoAmIResponse)
	err := c.cc.Invoke(ctx, Presence_AuthWhoAmI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceClient) AuthLogout(ctx context.Context, in *AuthLogoutRequest, opts ...grpc.CallOption) (*AuthLogoutResponse, error) {
	out := new(AuthLogoutResponse)
	err := c.cc.Invoke(ctx, Presence_AuthLogout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceClient) SessionGet(ctx context.Context, in *SessionGetRequest, opts ...grpc.CallOption) (*SessionGetResponse, error) {
	out := new(SessionGetResponse)
	err := c.cc.Invoke(ctx, Presence_SessionGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceClient) SessionPut(ctx context.Context, in *SessionPutRequest, opts ...grpc.CallOption) (*SessionPutResponse, error) {
	out := new(SessionPutResponse)
	err := c.cc.Invoke(ctx, Presence_SessionPut_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceClient) SessionDelete(ctx context.Context, in *SessionDeleteRequest, opts ...grpc.CallOption) (*SessionDeleteResponse, error) {
	out := new(SessionDeleteResponse)
	err := c.cc.Invoke(ctx, Presence_SessionDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresenceServer is the server API for Presence service.
// All implementations must embed UnimplementedPresenceServer
// for forward compatibility
type PresenceServer interface {
	// AuthCreateHandoffFlow creates a handoff flow with a given identifier and
	// password.
	AuthCreateHandoffFlow(context.Context, *AuthCreateHandoffFlowRequest) (*AuthCreateHandoffFlowResponse, error)
	// AuthCompleteHandoffFlow hands off a session to a gateway with a given token and
	// code.
	AuthCompleteHandoffFlow(context.Context, *AuthCompleteHandoffFlowRequest) (*AuthCompleteHandoffFlowResponse, error)
	// AuthLogin hands off a session to a gateway with a given token and
	// code.
	AuthLogin(context.Context, *AuthLoginRequest) (*AuthLoginResponse, error)
	// AuthRefreshLogin authenticates a gateway with a given identifier, password, and
	// token.
	AuthRefreshLogin(context.Context, *AuthRefreshLoginRequest) (*AuthRefreshLoginResponse, error)
	// AuthWhoAmI returns the session associated with a given token.
	AuthWhoAmI(context.Context, *AuthWhoAmIRequest) (*AuthWhoAmIResponse, error)
	// AuthLogout logs out a session with a given code.
	AuthLogout(context.Context, *AuthLogoutRequest) (*AuthLogoutResponse, error)
	// SessionGet gets a session with a given code.
	SessionGet(context.Context, *SessionGetRequest) (*SessionGetResponse, error)
	// SessionPut creates a session with a given identifier and token.
	SessionPut(context.Context, *SessionPutRequest) (*SessionPutResponse, error)
	// SessionDelete Deletes a session with a given code.
	SessionDelete(context.Context, *SessionDeleteRequest) (*SessionDeleteResponse, error)
	mustEmbedUnimplementedPresenceServer()
}

// UnimplementedPresenceServer must be embedded to have forward compatible implementations.
type UnimplementedPresenceServer struct {
}

func (UnimplementedPresenceServer) AuthCreateHandoffFlow(context.Context, *AuthCreateHandoffFlowRequest) (*AuthCreateHandoffFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthCreateHandoffFlow not implemented")
}
func (UnimplementedPresenceServer) AuthCompleteHandoffFlow(context.Context, *AuthCompleteHandoffFlowRequest) (*AuthCompleteHandoffFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthCompleteHandoffFlow not implemented")
}
func (UnimplementedPresenceServer) AuthLogin(context.Context, *AuthLoginRequest) (*AuthLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthLogin not implemented")
}
func (UnimplementedPresenceServer) AuthRefreshLogin(context.Context, *AuthRefreshLoginRequest) (*AuthRefreshLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthRefreshLogin not implemented")
}
func (UnimplementedPresenceServer) AuthWhoAmI(context.Context, *AuthWhoAmIRequest) (*AuthWhoAmIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthWhoAmI not implemented")
}
func (UnimplementedPresenceServer) AuthLogout(context.Context, *AuthLogoutRequest) (*AuthLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthLogout not implemented")
}
func (UnimplementedPresenceServer) SessionGet(context.Context, *SessionGetRequest) (*SessionGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionGet not implemented")
}
func (UnimplementedPresenceServer) SessionPut(context.Context, *SessionPutRequest) (*SessionPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionPut not implemented")
}
func (UnimplementedPresenceServer) SessionDelete(context.Context, *SessionDeleteRequest) (*SessionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionDelete not implemented")
}
func (UnimplementedPresenceServer) mustEmbedUnimplementedPresenceServer() {}

// UnsafePresenceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresenceServer will
// result in compilation errors.
type UnsafePresenceServer interface {
	mustEmbedUnimplementedPresenceServer()
}

func RegisterPresenceServer(s grpc.ServiceRegistrar, srv PresenceServer) {
	s.RegisterService(&Presence_ServiceDesc, srv)
}

func _Presence_AuthCreateHandoffFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCreateHandoffFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).AuthCreateHandoffFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_AuthCreateHandoffFlow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).AuthCreateHandoffFlow(ctx, req.(*AuthCreateHandoffFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Presence_AuthCompleteHandoffFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCompleteHandoffFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).AuthCompleteHandoffFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_AuthCompleteHandoffFlow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).AuthCompleteHandoffFlow(ctx, req.(*AuthCompleteHandoffFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Presence_AuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).AuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_AuthLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).AuthLogin(ctx, req.(*AuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Presence_AuthRefreshLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRefreshLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).AuthRefreshLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_AuthRefreshLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).AuthRefreshLogin(ctx, req.(*AuthRefreshLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Presence_AuthWhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthWhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).AuthWhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_AuthWhoAmI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).AuthWhoAmI(ctx, req.(*AuthWhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Presence_AuthLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).AuthLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_AuthLogout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).AuthLogout(ctx, req.(*AuthLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Presence_SessionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).SessionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_SessionGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).SessionGet(ctx, req.(*SessionGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Presence_SessionPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).SessionPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_SessionPut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).SessionPut(ctx, req.(*SessionPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Presence_SessionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).SessionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Presence_SessionDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).SessionDelete(ctx, req.(*SessionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Presence_ServiceDesc is the grpc.ServiceDesc for Presence service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Presence_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shikanime.elkia.fleet.v1alpha1.Presence",
	HandlerType: (*PresenceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthCreateHandoffFlow",
			Handler:    _Presence_AuthCreateHandoffFlow_Handler,
		},
		{
			MethodName: "AuthCompleteHandoffFlow",
			Handler:    _Presence_AuthCompleteHandoffFlow_Handler,
		},
		{
			MethodName: "AuthLogin",
			Handler:    _Presence_AuthLogin_Handler,
		},
		{
			MethodName: "AuthRefreshLogin",
			Handler:    _Presence_AuthRefreshLogin_Handler,
		},
		{
			MethodName: "AuthWhoAmI",
			Handler:    _Presence_AuthWhoAmI_Handler,
		},
		{
			MethodName: "AuthLogout",
			Handler:    _Presence_AuthLogout_Handler,
		},
		{
			MethodName: "SessionGet",
			Handler:    _Presence_SessionGet_Handler,
		},
		{
			MethodName: "SessionPut",
			Handler:    _Presence_SessionPut_Handler,
		},
		{
			MethodName: "SessionDelete",
			Handler:    _Presence_SessionDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/fleet/v1alpha1/fleet.proto",
}
